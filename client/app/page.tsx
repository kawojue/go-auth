import { Button } from "@/components/ui/button"
import Link from "next/link"

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <p>Go Auth</p>
      <Button>
        <Link href="/signup">
          Get Started
        </Link>
      </Button>
    </main>
  )
}
