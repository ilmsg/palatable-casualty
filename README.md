# palatable-casualty
gorm with role authorize


<user> can <action> <authorize> 
user a1 can create project
user a2 can read, update project

project
- members
  - project owner (Create, Read, Update, Delete)
  - member (Read, Update)

- user
  - email
  - password
- role
  - project owner
  - member
- authorize
  - create
  - read
  - update
  - delete
- project
  - title
  - tasks
  - members
    - user_id
    - role
- task
  - title

---
### Project Owner
||Create|Read|Update|Delete|
|---|---|---|---|---|
|Project|X|X|X|X|X|
|Task|X|X|X|X|X|

---

### Modulator
||Create|Read|Update|Delete|
|---|---|---|---|---|
|Project||X||||
|Task|X|X|X|X|X|

---

### Member
||Create|Read|Update|Delete|
|---|---|---|---|---|
|Project||X||||
|Task||X|X|||

---

### Viewer
||Create|Read|Update|Delete|
|---|---|---|---|---|
|Project||X||||
|Task||X||||
