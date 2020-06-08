# include <stdio.h>
# include <windows.h>
# include <stddef.h>
# include "resource.h"//������Դͷ�ļ�,<>����������·����ͷ�ļ�,""��������·��ͷ�ļ�
#undef UNICODE
void FreeResouce(UINT name,char *type,char * avename);

void FreeResouce(UINT name, char* type, char* savename)
{
	FILE* fp = NULL;
	HRSRC hSRC = FindResource(NULL,MAKEINTRESOURCE(name), type);
	if (NULL == hSRC)
		return;
	DWORD size = sizeof(NULL, hSRC);
	if (size <= 0)
		return;
	HGLOBAL hGlobal = LoadResource(NULL, hSRC);
	if (NULL == hGlobal)
		return;
	LPVOID lpvoid = LockResource(hGlobal);
	if (NULL == lpvoid)
		return;
	fopen_s(&fp, savename, "wb+");
	if (NULL == fp)
		return;
	fwrite(lpvoid, sizeof(char), size, fp);
	fclose(fp);

}
void main() {
	FreeResouce(IDR_VIRS2,"virs","free.exe");//IDR_VIRS2��resource.h���Զ�����,��������
}
