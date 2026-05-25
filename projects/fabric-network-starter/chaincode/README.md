# Smart LMS Chaincode

Go chaincode mẫu cho project `fabric-network-starter`.

## Functions

- `Init`
- `CreateCourse(id, title, instructor, description)`
- `ReadCourse(id)`
- `UpdateCourse(id, title, instructor, description, status)`
- `DeleteCourse(id)`
- `GetAllCourses`
- `CourseExists(id)`

## Ví dụ invoke/query

```json
{"Args":["CreateCourse","course-002","Blockchain 101","teacher1","Intro course"]}
```

```json
{"Args":["ReadCourse","course-002"]}
```

```json
{"Args":["GetAllCourses"]}
```
