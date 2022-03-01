module "user" {
    operation "getUserList" {
        description = "获取用户列表"
        command = <<EOF
        SELECT User {
            nickname, email
        }
        OFFSET {{.offset}}
        LIMIT {{.limit}}
        EOF
    }
    operation "getUserByID" {

    }
    operation "addUser" {

    }
    operation "updateUserByID" {

    }
    operation "deleteUserByID" {

    }
}