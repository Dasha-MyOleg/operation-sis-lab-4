# Dasha-MyOleg-operation-sis-lab-4

# Звіт до лабораторної роботи №3
**Тема:** Алгоритми заміни сторінок у віртуальній пам'яті  
**Дисципліна:** Системне програмне забезпечення (Операційні системи)  
**Студентка:** Дар'я Іванівна  
**Група:** ІМ-24  

---

## Мета роботи  
- Ознайомлення з механізмом роботи віртуальної пам'яті операційної системи.  
- Розробка програмної моделі для реалізації алгоритмів заміни сторінок.  
- Аналіз ефективності роботи алгоритмів на основі експериментальних даних.

---

##Приклад виконання команд.

```C:\Users\Даша\OneDrive\Документы\GitHub\operation-sis-lab-4>go run main.go
System initialization...
Creation of file system...
File system created
System is ready to work!
Create file: file1 | Size: 0
Create file: file2 | Size: 0
Hard links of current directory:
Name: file1      Size: 0
Name: file2      Size: 0
File file1 opened with descriptor ID: 0
Error: Write exceeds file size, consider truncating the file first
Error: Offset exceeds file size
Error: Read size exceeds file size
Closing file descriptor 0
Extending file: file1 to new size: 50
Hard links of current directory:
Name: file1      Size: 50
Name: file2      Size: 0
Delete file: file1
File file1 deleted successfully
Hard links of current directory:
Name: file2      Size: 0```
