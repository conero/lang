# 2018年12月19日 星期三
# python 类
# >= 3.6
# CA -> CB - CC



class CA:
    name = "CA"         # 类变量，可被所有实例共享
    counter = 0
    counterQue = []
    __number = 0
    __equalQue = []

    def __init__(self):
        self.counter += 1
        self.counterQue.append(1)

    def add(self, det=1):
        self.__number = self.__number + det
        self.__equalQue.append(' + '+str(det))
        return self

    def subtract(self, det=1):
        self.__number = self.__number - det
        self.__equalQue.append(' - '+str(det))
        return self

    def multiply(self, det=1):
        self.__number = self.__number * det    
        self.__equalQue.append(' * '+str(det))  
        return self

    def divide(self, det=1):
        self.__number = self.__number / det
        self.__equalQue.append(' / '+str(det))  
        return self

    def getNumber(self): 
        equal = '0' + ''.join(self.__equalQue) + ' = ' + str(self.__number)
        return self.__number, equal  
    
    def restar(self):
        __number = 0
        self.__equalQue = []
        return self

# 继承 CA -> CB
class CB(CA):
    name = "CB"
    queue = []

    def __init__(self, *arg):        
        self.queue = arg


# 继承 CB -> CC
class CC(CB):
    name = "CC"
    inniDick = {}


    def __init__(self, **cdick):
        self.inniDick = cdick
        self.__privateCode = '私有变量'

    # 静态变量
    @staticmethod
    def staticBase():
        print('CC 类的静态变量')
        pass



class TestConcole:
    pass
    def tCa(self):
        # test 1
        aIst = CA()
        aIst.add().add(50).multiply(2).divide(3)

        print(aIst.getNumber(), aIst.counterQue)

        # test 2
        aIst2 = CA()
        print(aIst2.getNumber(), aIst.counterQue)

        # test 3
        CC.staticBase()

        pass

    def __init__(self):
        self.tCa()       


TestConcole()