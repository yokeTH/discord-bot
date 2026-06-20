/*
  Warnings:

  - Added the required column `created_by` to the `watchlist` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "watchlist" ADD COLUMN     "created_by" INT8 NOT NULL;
