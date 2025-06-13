"use client"

import { useEffect, useState } from "react"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"

interface User {
  id: string
  name: string
  email: string
  created_at: string
  updated_at: string
}

export function UserList() {
  const [users, setUsers] = useState<User[]>([])
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/v1/users")
        if (!response.ok) {
          throw new Error("ユーザー一覧の取得に失敗しました")
        }
        const data = await response.json()
        setUsers(data)
      } catch (error) {
        setError(error instanceof Error ? error.message : "エラーが発生しました")
      }
    }

    fetchUsers()
  }, [])

  if (error) {
    return <div className="text-red-500">{error}</div>
  }

  return (
    <Card className="w-[600px]">
      <CardHeader>
        <CardTitle>ユーザー一覧</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="space-y-4">
          {users.map((user) => (
            <div
              key={user.id}
              className="flex justify-between items-center p-4 border rounded-lg"
            >
              <div>
                <div className="font-medium">{user.name}</div>
                <div className="text-sm text-gray-500">{user.email}</div>
              </div>
              <div className="text-sm text-gray-500">
                作成日: {new Date(user.created_at).toLocaleDateString()}
              </div>
            </div>
          ))}
        </div>
      </CardContent>
    </Card>
  )
}
