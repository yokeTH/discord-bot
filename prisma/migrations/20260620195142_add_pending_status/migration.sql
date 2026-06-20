/*
  Warnings:

  - Added the required column `created_by` to the `pending_delete` table without a default value. This is not possible if the table is not empty.

*/
-- CreateEnum
CREATE TYPE "PendingStatus" AS ENUM ('pending', 'confirmed', 'cancelled');

-- AlterTable
ALTER TABLE "pending_delete" ADD COLUMN     "created_by" INT8 NOT NULL;
ALTER TABLE "pending_delete" ADD COLUMN     "status" "PendingStatus" NOT NULL DEFAULT 'pending';
