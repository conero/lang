# 多[进程]学习
# 2019年9月27日 星期五
# python 3.7+
import os
import multiprocessing as mp
import time

'''
    每个进程都有不同的 pid

'''


def worker():
    time.sleep(1)
    print(f'worker => time: {time.time()}')
    time.sleep(1)
    print(f'worker => pid: {os.getpid()}')


# 继承式进程
class MyTask(mp.Process):
    def __init__(self, name):
        super().__init__()
        self.name = name

    def run(self):
        time.sleep(1)
        print(f'{self.name}: pid => {os.getpid()}')
        time.sleep(1)
        print(f'{self.name}: time => {time.time()}')


if __name__ == '__main__':
    # 主进程
    print('系统进入主进程：')
    pQue = []

    # 线程 1
    wTh1 = mp.Process(target=worker)
    wTh1.start()
    pQue.append(wTh1)

    # 线程2
    wTh2 = mp.Process(target=worker)
    wTh2.start()
    pQue.append(wTh2)

    # 继承式进程
    p1 = MyTask('p1')
    p1.start()
    pQue.append(p1)

    p2 = MyTask('p2')
    p2.start()
    pQue.append(p2)

    # 等待所有线程完成
    for p in pQue:
        p.join()
    print(f'主进程已退出: {os.getpid()}')
