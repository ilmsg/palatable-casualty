# palatable-casualty
gorm with role authorize


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
