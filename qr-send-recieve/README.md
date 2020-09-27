# ������ ��������, ��������, ��������� � ������������� QR ����

(�) softland (softlandia@gmail.com)

2020-09-27

## qr-server

������� :8080. ������������ post ��������� �� ������ /qr

body ������ ��������� json �������

```bash
{
	`json:"employeeID"` // uint64
	`json:"officeID"`   // uint64
	`json:"date"`       // 2009-11-10 23:00:00 +0000 UTC
}
```

� �������� ������ ������������ png ����������� QR ���� � �������������� json

�������� �� �����


```bash
{
	"employeeID":1,
	"officeID":911,
	"date":2009-11-10 23:00:00 +0000 UTC
}
```

�� ������ ����� QR �� ������� 

```bash
{"employeeID":1,"officeID":911,"date":2009-11-10 23:00:00 +0000 UTC}
```

## qr-tester

������� :8081. ������������ get ��������� �� ������ /getQR  
��������� json � ���������� post ������ � qr-server  
�������� ������� png � QR �����, ���������� ��� � ������� ������ � log