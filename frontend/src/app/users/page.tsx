import { CreateUserForm } from "@/components/user/create-user-form"
import { UserList } from "@/components/user/user-list"

export default function UsersPage() {
  return (
    <div className="container mx-auto py-8">
      <h1 className="text-3xl font-bold mb-8">ユーザー管理</h1>
      <div className="flex gap-8">
        <CreateUserForm />
        <UserList />
      </div>
    </div>
  )
}
