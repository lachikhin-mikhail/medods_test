CREATE TABLE "users" (
    "id" INTEGER,
    "username" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    PRIMARY KEY ("id" AUTOINCREMENT)
);

CREATE TABLE "refresh_tokens" (
    "id" INTEGER,
    "user_id" INTEGER NOT NULL,
    "token" TEXT NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT),
    FOREIGN KEY("user_id")
    
)