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
        setUsers(data.users || [])
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
    <Card className="w-[600px] mx-auto p-4">
      <CardHeader>
        <CardTitle>ユーザー一覧</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="flex flex-col gap-3 px-2 py-2">
          {users.length === 0 ? (
            <div className="text-center text-gray-500 py-8">
              ユーザーが存在しません
            </div>
          ) : (
            users.map((user) => (
              <div
                key={user.id}
                className="flex justify-between items-center px-4 py-3 bg-white rounded-lg border border-gray-100 shadow-sm hover:bg-gray-50 transition-colors"
              >
                <div>
                  <div className="font-bold text-base text-gray-900 mb-0.5">{user.name}</div>
                  <div className="text-xs text-gray-500 mb-0.5">{user.email}</div>
                </div>
                <div className="text-xs text-gray-400 text-right min-w-[100px]">
                  作成日:<br />{new Date(user.created_at).toLocaleDateString()}
                </div>
              </div>
            ))
          )}
        </div>
      </CardContent>
    </Card>
  )
}
