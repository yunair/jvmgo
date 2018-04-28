public class BubbleSortTest {
    public static void main(String[] args){
        int[] arr = {22, 84, 77, 11, 95, 9, 78, 56, 34, 13, 65, 43, 10, 34, 92};
        bubbleSort(arr);
        printArray(arr);
    }

    private static void bubbleSort(int[] arr) {
        boolean swapped = true;
        int j = 0;
        int tmp;
        while(swapped) {
            swapped = false;
            j ++;
            for (int i = 0; i < arr.length - 1; i++) {
                if (arr[i] > arr[i+1]){
                    tmp = arr[i];
                    arr[i] = arr[i+1];
                    arr[i+1] = tmp;
                    swapped = true;
                }
            }
        }
    }

    private static void printArray(int[] arr){
        for (int i : arr){
            System.out.println(i);
        }

    }
}