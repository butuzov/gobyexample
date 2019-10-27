$ go run directories.go
Перегляд subdir/parent
  child true
  file2 false
  file3 false
Перегляд subdir/parent/child
  file4 false
Відвідуємо subdir
  subdir true
  subdir/file1 false
  subdir/parent true
  subdir/parent/child true
  subdir/parent/child/file4 false
  subdir/parent/file2 false
  subdir/parent/file3 false
