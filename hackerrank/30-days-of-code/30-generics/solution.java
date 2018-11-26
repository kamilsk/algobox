class Printer <T> {
    public void printArray(T[] list) {
        for (T el : list) {
            System.out.println(el);
        }
    }
}
