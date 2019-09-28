# 多[线程]学习
# 2019年9月27日 星期五
# python 3.7+
import os
import threading
import time


def worker():
    time.sleep(2)
    print(time.time())
    print(os.getpid())


# 线程-任务1
class Task1(threading.Thread):
    def __init__(self, threadId, name):
        # super().__init__()
        threading.Thread.__init__(self)
        self.innerDick = {
            'id': threadId,
            'name': name
        }

    # 运行线程
    def run(self):
        print(f"线程运行: {self.innerDick['id']}")


if __name__ == '__main__':
    # 主进程
    print('系统进入主进程：')
    # 线程列表
    threads = []

    # 线程 1
    wTh1 = threading.Thread(target=worker)
    wTh1.start()

    # 线程2
    wTh2 = threading.Thread(target=worker)
    wTh2.start()

    print(os.getpid())

    # 添加线程到线程列表
    threads.append(wTh1)
    threads.append(wTh2)

    # 继承的线程运行
    tsk1 = Task1('class-1', '运行任务一')
    tsk1.start()
    threads.append(tsk1)

    tsk2 = Task1('class-2', '运行任务二')
    tsk2.start()
    threads.append(tsk2)

    # 等待所有线程完成
    for t in threads:
        t.join()
    print("退出主线程")
