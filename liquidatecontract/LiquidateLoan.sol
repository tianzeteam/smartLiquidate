// SPDX-License-Identifier: agpl-3.0
pragma solidity 0.6.12;
pragma experimental ABIEncoderV2;
import { FlashLoanReceiverBase } from "./FlashLoanReceiverBase.sol";
import { ILendingPool, ILendingPoolAddressesProvider, IERC20 } from "./Interfaces.sol";
import { SafeMath } from "./Libraries.sol";
import "./Ownable.sol";

import "./IUniswapV2Router02.sol";


/*
* A contract that liquidates an aave loan using a flash loan:
*
*   call executeFlashLoans() to begin the liquidation
*
*/
interface IPriceOracleGetter {
  function getAssetPrice(address asset) external view returns (uint256);
}
interface IProtocolDataProvider {


  function getUserReserveData(address asset, address user)
    external
    view
    returns (
      uint256 currentATokenBalance,
      uint256 currentStableDebt,
      uint256 currentVariableDebt,
      uint256 principalStableDebt,
      uint256 scaledVariableDebt,
      uint256 stableBorrowRate,
      uint256 liquidityRate,
      uint40 stableRateLastUpdated,
      bool usageAsCollateralEnabled
    );
}

contract LiquidateLoan is FlashLoanReceiverBase, Ownable {

    IUniswapV2Router02 public uniswapV2Router;

    IPriceOracleGetter public priceOracleGetter;
    IProtocolDataProvider public protocolDataProvider ;

    using SafeMath for uint256;
    event ErrorHandled(string stringFailure);
   
    // intantiate lending pool addresses provider and get lending pool address
    constructor(ILendingPoolAddressesProvider _addressProvider, IUniswapV2Router02 _uniswapV2Router,address _protocolDataProvider,address _priceOracleGetter) FlashLoanReceiverBase(_addressProvider) public {
        // instantiate UniswapV2 Router02
        uniswapV2Router = IUniswapV2Router02(address(_uniswapV2Router)); 

        protocolDataProvider = IProtocolDataProvider(address(_protocolDataProvider));

        priceOracleGetter = IPriceOracleGetter(address(_priceOracleGetter));
    }
    
    function setPriceOralce(address _priceOracleGetter) external onlyOwner{
         priceOracleGetter = IPriceOracleGetter(_priceOracleGetter);
    }

    function setProtocolDataProvider(address _protocolDataProvider) external onlyOwner {
           protocolDataProvider = IProtocolDataProvider(_protocolDataProvider);
    }
    struct TokenInfo {
        address tokenAddress;
        string  tokenName;
        uint256 decimals;
    }
    
    TokenInfo[]   public    tokenInfos ;


     function setToken(address  _tokenAddress,string memory _tokenName,uint256 decimals) external onlyOwner{
       tokenInfos.push(TokenInfo({tokenAddress: _tokenAddress, tokenName: _tokenName, decimals:decimals}));

     }
     
    function setTokens(address[] memory _tokenAddress,string[] memory _tokenName,uint256[] memory decimals) external onlyOwner{
        uint length = _tokenAddress.length;
        for(uint i=0;i<length;i++){
              tokenInfos.push(TokenInfo({tokenAddress: _tokenAddress[i], tokenName: _tokenName[i], decimals:decimals[i]}));
        }
 
     }
     
    function removeToken(uint tokenIndex) public onlyOwner{
        delete tokenInfos[tokenIndex];
    }

     struct CollateralReserve {
         uint256 currentATokenBalance ;
         string symbol;
         address underlyingAsset;
         uint256 decimals;
         bool usageAsCollateralEnabled;
         uint256 priceInEth;
         bool notNull;
        
     }

     struct BorrowReserve {
         uint256 currentTotalDebt ;
         string symbol;
         address underlyingAsset;
         uint256 decimals;
         bool usageAsCollateralEnabled;
         uint256 priceInEth;
         bool notNull;
     }

     function getLastUserAssetData(address  _user) external view returns( CollateralReserve[] memory ,  BorrowReserve[] memory){
         
        CollateralReserve[] memory _collateralReserves =  new CollateralReserve[](6);
        BorrowReserve[] memory _borrowReserves =  new BorrowReserve[](6);
        uint8 collateralLastIndex =0;
        uint8 borrowLastIndex = 0;
        {
            address   liquidateUser = _user;
            for(uint256 i=0;i<18;i++){
                    TokenInfo memory tokenInfo =tokenInfos[i];
                    address tokenAddress = tokenInfo.tokenAddress;
                    uint  decimals = tokenInfo.decimals;
                    string memory tokenName = tokenInfo.tokenName ;

                    (uint256 _currentATokenBalance,
                     uint256 _currentStableDebt,
                     uint256 _currentVariableDebt,
                     ,
                     ,
                     ,
                     ,
                     ,
                     bool _usageAsCollateralEnabled
                    ) = protocolDataProvider.getUserReserveData(tokenAddress, liquidateUser);

                    if (_currentATokenBalance > 0) {
                        uint256 _priceInEth = priceOracleGetter.getAssetPrice(tokenAddress);
                        _collateralReserves[collateralLastIndex]=CollateralReserve({currentATokenBalance:_currentATokenBalance,symbol:tokenName,underlyingAsset: tokenAddress,decimals:decimals,usageAsCollateralEnabled:_usageAsCollateralEnabled,priceInEth:_priceInEth,notNull:true});
                        collateralLastIndex ++ ;
                    }
                    
                    if (_currentStableDebt > 0 || _currentVariableDebt > 0){
                        uint256 _priceInEth = priceOracleGetter.getAssetPrice(tokenAddress);
                        uint256 _currentTotalDebt = _currentStableDebt.add(_currentVariableDebt);
                        _borrowReserves[borrowLastIndex]= BorrowReserve({currentTotalDebt:_currentTotalDebt,symbol:tokenName,underlyingAsset: tokenAddress,decimals:decimals,usageAsCollateralEnabled:_usageAsCollateralEnabled,priceInEth:_priceInEth,notNull:true});
                        borrowLastIndex ++ ;
                    } 
                    
                }
        }
  
       return (_collateralReserves , _borrowReserves);
     }

     function getLastUserAssetDataExtent(address  _user) external view returns( CollateralReserve[] memory ,  BorrowReserve[] memory){
         
        uint256  length = tokenInfos.length;
        CollateralReserve[] memory _collateralReserves =  new CollateralReserve[](6);
        BorrowReserve[] memory _borrowReserves =  new BorrowReserve[](6);
        uint8 collateralLastIndex =0;
        uint8 borrowLastIndex = 0;
        {
            address   liquidateUser = _user;
            for(uint256 i=18;i<length;i++){
                    TokenInfo memory tokenInfo =tokenInfos[i];
                    address tokenAddress = tokenInfo.tokenAddress;
                    uint  decimals = tokenInfo.decimals;
                    string memory tokenName = tokenInfo.tokenName ;

                    (uint256 _currentATokenBalance,
                     uint256 _currentStableDebt,
                     uint256 _currentVariableDebt,
                     ,
                     ,
                     ,
                     ,
                     ,
                     bool _usageAsCollateralEnabled
                    ) = protocolDataProvider.getUserReserveData(tokenAddress, liquidateUser);

                    if (_currentATokenBalance > 0) {
                        uint256 _priceInEth = priceOracleGetter.getAssetPrice(tokenAddress);
                        _collateralReserves[collateralLastIndex]=CollateralReserve({currentATokenBalance:_currentATokenBalance,symbol:tokenName,underlyingAsset: tokenAddress,decimals:decimals,usageAsCollateralEnabled:_usageAsCollateralEnabled,priceInEth:_priceInEth,notNull:true});
                        collateralLastIndex ++ ;
                    }
                    
                    if (_currentStableDebt > 0 || _currentVariableDebt > 0){

                        uint256 _priceInEth = priceOracleGetter.getAssetPrice(tokenAddress);
                        uint256 _currentTotalDebt = _currentStableDebt.add(_currentVariableDebt);
                        _borrowReserves[borrowLastIndex]= BorrowReserve({currentTotalDebt:_currentTotalDebt,symbol:tokenName,underlyingAsset: tokenAddress,decimals:decimals,usageAsCollateralEnabled:_usageAsCollateralEnabled,priceInEth:_priceInEth,notNull:true});
                        borrowLastIndex ++ ;
                    } 
                    
                }
        }
  
       return (_collateralReserves , _borrowReserves);
     }
    /**
        This function is called after your contract has received the flash loaned amount
     */
    function executeOperation(
        address[] calldata _assets,
        uint256[] calldata _amounts,
        uint256[] calldata _premiums,
        address ,
        bytes calldata params
    )
        external
        override
        returns (bool)
    {
        address  borrowAsset = _assets[0];
        uint256 liquidateAmount = _amounts[0];
        uint256 premium = _premiums[0];
        //collateral  the address of the token that we will be compensated in
        //userToLiquidate - id of the user to liquidate
        //amountOutMin - minimum amount of asset paid when swapping collateral

        (address collateral, address userToLiquidate, uint256 amountOutMin, address[] memory swapPath,uint256 gasCost) = abi.decode(params, (address, address, uint256, address[],uint256));

        //liquidate unhealthy loan
        liquidateLoan(collateral, borrowAsset, userToLiquidate, liquidateAmount, false);

        //swap collateral from liquidate back to asset from flashloan to pay it off
        uint256 length = swapPath.length;
        if(swapPath[0] != swapPath[length-1]){
            swapToBarrowedAsset(collateral,amountOutMin, swapPath);
        }
        

        //Pay to owner the balance after fees
        uint256 profit = calcProfits(IERC20(borrowAsset).balanceOf(address(this)),liquidateAmount,premium,gasCost);

        require(profit > 0 , "No profit");
        IERC20(borrowAsset).transfer(owner(), profit);


        // Approve the LendingPool contract allowance to *pull* the owed amount
        // i.e. AAVE V2's way of repaying the flash loan
        uint repayAmount = liquidateAmount.add(premium).add(1);
        IERC20(borrowAsset).approve(address(_lendingPool), repayAmount);

        return true;
    }

    //calculate profits after paying back loan & fees
    function calcProfits(uint256 _balance, uint256 _loanAmount, uint256 _loanFee,uint256 _gasCost)
        pure
        private
        returns(uint256)
    {
        return _balance.sub(_loanAmount.add(_loanFee).add(_gasCost),"no profits to return");
    }

    function liquidateLoan(address _collateral, address _liquidate_asset, address _userToLiquidate, uint256 _amount, bool _receiveaToken) public {

        require(IERC20(_liquidate_asset).approve(address(_lendingPool), _amount), "Approval error");

        _lendingPool.liquidationCall(_collateral,_liquidate_asset, _userToLiquidate, _amount, _receiveaToken);
    }


    //assumes the balance of the token is on the contract
    function swapToBarrowedAsset(address asset_from, uint amountOutMin, address[] memory swapPath ) public {

        // setting deadline to avoid scenario where miners hang onto it and execute at a more profitable time
        uint deadline = block.timestamp + 120; // 2 minutes

        uint256 amountIn = IERC20(asset_from).balanceOf(address(this));

        // grant uniswap access to your token
        IERC20(asset_from).approve(address(uniswapV2Router), uint256(-1));

        // Trade 1: Execute swap from asset_from into designated ERC20 (asset_to) token on UniswapV2
        try uniswapV2Router.swapExactTokensForTokens(
            amountIn,
            amountOutMin,
            swapPath,
            address(this),
            deadline
        ){
        }
        catch Error(string memory reason)
        {
            emit ErrorHandled(reason);
        }
        catch
        {

        }

    }


    /*
    * This function is manually called to commence the flash loans sequence
    * to make executing a liquidation  flexible calculations are done outside of the contract and sent via parameters here
    * _assetToLiquidate - the token address of the asset that will be liquidated
    * _flashAmt - flash loan amount (number of tokens) which is exactly the amount that will be liquidated
    * _collateral - the token address of the collateral. This is the token that will be received after liquidating loans
    * _userToLiquidate - user ID of the loan that will be liquidated
    * _amountOutMin - when using uniswap this is used to make sure the swap returns a minimum number of tokens, or will revert
    * _swapPath - the path that uniswap will use to swap tokens back to original tokens

    */
    function executeFlashLoans(address _assetToLiquidate, uint256 _flashAmt, address _collateral, address _userToLiquidate, uint256 _amountOutMin, address[] memory _swapPath,uint256 _gasCost) public onlyOwner {

        // the various assets to be flashed
        address[] memory assets = new address[](1);
        assets[0] = _assetToLiquidate;

        // the amount to be flashed for each asset
        uint256[] memory amounts = new uint256[](1);
        amounts[0] = _flashAmt;

        // 0 = no debt, 1 = stable, 2 = variable
        uint256[] memory modes = new uint256[](1);
        modes[0] = 0;

        //only for testing. must remove
        // passing these params to executeOperation so that they can be used to liquidate the loan and perform the swap
        bytes memory params = abi.encode(_collateral, _userToLiquidate, _amountOutMin, _swapPath,_gasCost);

        _lendingPool.flashLoan(
            address(this),
            assets,
            amounts,
            modes,
            address(this),
            params,
            0
        );
    }



}
