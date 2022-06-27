//SPDX-License-Identifier: LICENSED

// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.7.0;
import "@openzeppelin/contracts/access/Ownable.sol";
import "./TransferHelper.sol";
import "./SafeMath.sol";
import "./ERC20Interface.sol";

contract BiddingWar is Ownable {
    using SafeMath for uint256;
    address public NRG;
    uint256 public bidEndTime = 60 minutes;
    uint256 public maxBidInterval = 10 minutes;
    uint256 public commission = 50; // 50 = 5%
    bool public startBidding = false;
    uint256 public startTime;
    uint256 public curBidAmount;
    uint256 public curBidTime;
    address public lastBidder;
    uint256 public gameIndex = 0;
    modifier gameRunning() {
        require(checkGameRunning(), "game ended");
        _;
    }
    modifier gameEnded() {
        require(!checkGameRunning(), "game is running");
        _;
    }
    event GameStarted(uint256 gameIndex, uint256 startTime);
    event WinnerGetRewards(uint256 gameIndex, address winner, uint256 amount);
    event GameParamChanged(
        uint256 bidEndTime,
        uint256 maxBidInterval,
        uint256 commission
    );
    event NRGTokenAddressUpdated(address token);
    event Bid(
        uint256 ganeIndex,
        address bidder,
        uint256 bidAmount,
        uint256 bidTime
    );

    function checkGameRunning() public view returns (bool) {
        uint256 currentTime = block.timestamp;
        if (
            startBidding &&
            currentTime < startTime.add(bidEndTime) &&
            currentTime < curBidTime.add(maxBidInterval)
        ) return true;
        else return false;
    }

    function start() public onlyOwner gameEnded {
        //
        startTime = block.timestamp;
        curBidTime = startTime;
        curBidAmount = 0;
        startBidding = true;
        lastBidder = address(0);
        gameIndex = gameIndex + 1;
        emit GameStarted(gameIndex, startTime);
    }

    function rewardToWinner() public onlyOwner gameEnded {
        uint256 rewards = ERC20Interface(NRG).balanceOf(address(this));
        TransferHelper.safeTransfer(NRG, lastBidder, rewards);
        emit WinnerGetRewards(gameIndex, lastBidder, rewards);
        start();
    }

    function setTokenAddress(address token) public onlyOwner {
        NRG = token;
        emit NRGTokenAddressUpdated(token);
    }

    function setGameParams(
        uint256 _bidEndTime,
        uint256 _maxBidInterval,
        uint256 _commission
    ) public onlyOwner {
        bidEndTime = _bidEndTime;
        maxBidInterval = _maxBidInterval;
        commission = _commission;
        require(commission < 1000, "commission overflow");
        emit GameParamChanged(bidEndTime, maxBidInterval, commission);
    }

    function bid(uint256 bidAmount) external gameRunning {
        require(bidAmount > curBidAmount, "bid amount is less than before");
        lastBidder = msg.sender;
        TransferHelper.safeTransferFrom(
            NRG,
            lastBidder,
            address(this),
            bidAmount
        );
        // send commistion to owner
        uint256 commissionAmount = bidAmount.mul(commission).div(1000);
        TransferHelper.safeTransfer(NRG, owner(), commissionAmount);
        curBidTime = block.timestamp;
        curBidAmount = bidAmount;
        emit Bid(gameIndex, lastBidder, bidAmount, curBidTime);
    }
}
