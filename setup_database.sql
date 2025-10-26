-- Setup Database untuk MAMP MySQL
-- Jalankan script ini di phpMyAdmin atau MySQL CLI

CREATE DATABASE IF NOT EXISTS `golang-invoice` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `golang-invoice`;

-- Catatan: Table akan dibuat otomatis oleh GORM AutoMigrate
-- Jalankan aplikasi dengan: go run main.go
-- GORM akan membuat table: products, patients, invoices, invoice_items
