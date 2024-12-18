package java;

class FinalPrices {

    class Solution {
        public int[] finalPrices(int[] prices) {
            int[] answer = new int[prices.length];

            for (int i = 0; i < prices.length; i++) {
                for (int j = i; j < prices.length; j++) {
                    if (j > i && prices[j] <= prices[i]) {
                        answer[i] = prices[i] - prices[j];
                        break;
                    }
                    if (j == (prices.length - 1)) {
                        answer[i] = prices[i];
                    }
                }
            }

            return answer;
        }
    }
}