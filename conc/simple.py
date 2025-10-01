# 共享数据
import threading
import time
from concurrent.futures import ThreadPoolExecutor


shareCount = 0


# 创建锁对象
lock = threading.Lock()


def task_run(task_id):
    """
    线程函数
    """
    print(f"线程: {task_id}\r", end="")
    global shareCount
    lock.acquire()
    if task_id % 2 == 0:
        shareCount += 2
    else:
        shareCount -= 1
    lock.release()


def run_simple(count):
    """
    线程入口主函数
    """

    task_list = []
    for i in range(count):
        thread_el = threading.Thread(target=task_run, args=(i,))
        thread_el.start()
        task_list.append(thread_el)

    # 等待所有线程执行完毕
    for el in task_list:
        el.join()


def run_with_thread_pool(task_count, max_workers=100):
    """
    使用线程池处理大量任务
    :param task_count: 总任务数（可以非常大）
    :param max_workers: 最大并发线程数（根据系统性能调整）
    """
    # 创建线程池，指定最大线程数
    with ThreadPoolExecutor(max_workers=max_workers) as executor:
        # 提交所有任务（任务数量可无限大，线程池自动调度）
        executor.map(task_run, range(task_count))


def run_thread_simple():
    """
    简单的多线测试，线程数无法无限变大
    """
    print("--------- python 多线程数据测试 ---------")
    # count = 1_987_654

    start = time.perf_counter()
    count = 1_987_6
    run_simple(count)
    end = time.perf_counter()

    print()
    print(f"线程数：{count}, shareCount: {shareCount}，耗时: {end - start:.6f} 秒")


def run_thread_pool_test():
    print("--------- 线程池多任务测试 ---------")
    start = time.perf_counter()

    # 任务数量可以非常大（例如 100万），线程池会自动处理
    # 1_987_654 时程序可会卡死
    # task_count = 1_987_654
    task_count = 1_000_654
    # 最大并发线程数：根据系统调整（一般设为 CPU核心数*5 或 100-1000 不等）
    run_with_thread_pool(task_count, max_workers=500)

    end = time.perf_counter()
    print()  # 换行，避免被 \r 覆盖
    print(
        f"总任务数：{task_count}, shareCount: {shareCount}，耗时: {end - start:.6f} 秒"
    )


if __name__ == "__main__":
    run_thread_simple()
    print()

    # shareCount = 0
    # run_thread_pool_test()
