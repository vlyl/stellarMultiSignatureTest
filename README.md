# stellarMultiSignatureTest

test multi-signature of stellar network

description：
>构建这样一个事务：

>从账户A向账户Z发起一笔转账，金额随意

>该账户A需要A和B的签名，A和B的权重相等都为1，事务的中级阈值设置为2（转账属于中级操作）

>分别测试只A签名、只B签名、A和B签名、A,B和C签名时，事务是否成功
