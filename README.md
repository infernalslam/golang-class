# golang
### การ mocking data เพื่อทำการเทส functon
#### สิ่งที่ต้องเตรียมคือการประการตัวแปรแบบ struct เพื่อทำการแชร์ data ที่ mock ไปยังไฟล์เทสได้นั้นเอง


## จงทำให้ test ผ่าน
### โจทย์ที่ยกตัวอย่าง

You have been asked to write a program which counts from 3, printing each number on a new line (with a 1 second pause) and when it reaches zero it will print "Go!" and exit.



```
3
2
1
Go!
```

 * เริ่มต้นด้วยการสร้างไฟล์ main.go เขียน function main
 ```go
 package main

func main() {
    Countdown()
}
 ```

 * เริ่มต้นด้วยไฟล์เทสตาม code นี้
 ```go
 func TestCountdown(t *testing.T) {
    buffer := &bytes.Buffer{}

    Countdown(buffer)

    got := buffer.String()
    want := "3"

    if got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}
 ```

* รัน `go test .` result ที่ได้ ` undefined: Countdown`

```
เฉลยอยู่ใน file mocking-again
```
