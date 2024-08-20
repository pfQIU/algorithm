package code

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-100-liked
func MaxProfit(prices []int) int {
	pre, ans := prices[0], 0
	for i := range prices {
		ans = max(ans, prices[i]-pre)
		pre = min(pre, prices[i])
	}
	return ans
}
