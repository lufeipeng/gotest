# gotest
ѧϰGo��Project

//2015-10-19

//�ҵ�go��������
���в��������ķ���, ���ش���󣬽����Ӧ��Ŀ¼������mydoc, ����go test -v ����
��Ŀ¼��webserver.go �򵥵�webserver���Ժ���
main.go ��ں���, elipse����һ��main packageֻ����һ��main func, ��ʱ�޸�Ϊ��main_back

���õ���windows����eclipse + php plugin
���Ѳ��ԵĴ�����뵽github����, ����ͨ��github����Դ�������

//����godoc

godoc����ʵʱ��ˢ��, �޸�Դ�����doc ��ҳֱ��ˢ�¾Ϳ��Կ���

���е����� godoc -http:6060 
���Կ���goroot�����Լ���package��Ϣ, �����Լ�����Ĺ���
http://localhost:6060/pkg/github.com/lufeipeng/gotest/godoc/
��function�в���ʹ��fmt�����, go test ��ʱ��ᱨ��

//gotest
����Ŀ¼����mystrings package
go test ./mystrings -v -x
ע��godocĿ¼��example�ļ�example_test.go, ע�͵�output������ʵ��function�������Ӧ��, ����go test ʧ��

//oracle ����eclipse ctrl + �Ҽ���ת����
��Ҫ����golang tools, ���ԭ�������ص�ַʧЧ
go get -u golang/x/tools/cmd/oracle
��github.com\golang\tools ����tools ���뵽gopath�Ķ�Ӧ·��, ������go install ����


