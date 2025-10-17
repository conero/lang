package conc;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

// 参考：  https://www.cnblogs.com/jimoer/p/19083170    技术面:Java并发(线程同步、死锁、多线程编排)
public class Simple {
    public static void main(String[] args) {
        System.out.println("--------------------------- Java 简单的并发测试 ---------------------------");
        System.out.println();
        System.out.println();

        // 测试线程
        testSimpleThread();
    }

    // 测试线程
    private static void testSimpleThread() {
        System.out.println("--------------------------- 1. 线程创建 ---------------------------");

        // 1. 记录开始时间（纳秒）
        long startNano = System.nanoTime();
        int num = 1_987_654;
        int pool = 100;
        int countData = 0;
        try {
            countData = simplePool(num, pool);
        } catch (Exception e) {
            System.out.println("发生异常：" + e.getMessage());
        }
        // 3. 记录结束时间（纳秒）
        long endNano = System.nanoTime();
        System.out.println();
        System.out.println("执行完毕");
        System.out.println("任务数：" + num + "，计算值：" + countData + "，线程池：" + pool + ", 耗时："
                + ((endNano - startNano) / 1000_000_000.0) + " 秒");

    }

    // 启动线程池
    private static int simplePool(int num, int pool) throws InterruptedException {
        // ExecutorService executorService = Executors.newFixedThreadPool(pool);
        ExecutorService executorService = Executors.newScheduledThreadPool(pool);
        SharedData sharedData = new SharedData(0);

        for (int i = 0; i < num; i++) {
            SimpleTask simpleTask = new SimpleTask(i, sharedData);
            executorService.execute(simpleTask);
        }

        // 关闭线程池（不再接收新任务）
        executorService.shutdown();

        boolean allDone = executorService.awaitTermination(300, TimeUnit.SECONDS);
        // 5. 所有任务完成后，处理后续业务
        if (allDone) {
            System.out.println("所有任务已执行完毕，开始处理后续业务...");
            // 执行后续操作（如汇总结果、释放资源等）
        } else {
            System.out.println("等待超时，仍有任务未完成！");
        }

        return sharedData.getValue();
    }

    // 执行任务
    private static class SimpleTask extends Thread {
        private final int index;
        private final SharedData sharedData;

        public SimpleTask(int iptIndex, SharedData sharedData) {
            this.index = iptIndex;
            this.sharedData = sharedData;
        }

        @Override
        public void run() {
            // System.out.println(index + ": " +
            // Thread.currentThread().getName() + " is running");
            System.out.print(index + ": " + Thread.currentThread().getName() + " is running ...\r");
            if (this.index % 2 == 0) {
                sharedData.add(2);
            } else {
                sharedData.add(-1);
            }
        }
    }

    // 自定义共享数据类
    // 线程安全；加锁同步（适合复杂对象共享）
    private static class SharedData {
        private int value; // 需要共享的字段

        public SharedData(int initial) {
            this.value = initial;
        }

        // 同步方法：保证修改操作的线程安全
        public synchronized void add(int detData) {
            value += detData;
        }

        // 同步方法：保证读取操作的线程安全
        public synchronized int getValue() {
            return value;
        }
    }
}
