-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 02 Bulan Mei 2024 pada 05.00
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.0.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `crud_gin`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `barangs`
--

CREATE TABLE `barangs` (
  `id` int(11) NOT NULL,
  `product_name` varchar(255) NOT NULL,
  `category_id` int(20) NOT NULL,
  `unit` varchar(255) NOT NULL,
  `price` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `barangs`
--

INSERT INTO `barangs` (`id`, `product_name`, `category_id`, `unit`, `price`) VALUES
(1, 'kenjler', 1, 'kintil', 'kuntul'),
(3, 'makanan ringan', 2, 'cireng', 'limaribu'),
(4, 'makanan berat', 3, 'cilok', 'limaribu'),
(5, 'makanan enteng', 3, '', ''),
(6, 'makanan cihuy', 3, '', ''),
(8, 'makanan asoy', 7, 'cireng', 'seribu'),
(9, 'champ', 0, 'xyz', 'zyx'),
(10, 'kenzler', 0, 'abcd', 'dcba'),
(18, 'ejer', 0, 'dgeg', 'egeg'),
(19, 'sonice', 0, 'xxx', 'zzzz'),
(20, 'bintang-zero', 0, 'duapuluh', 'limaribu'),
(21, 'smirnof', 0, 'duapuluh', 'limaribu'),
(22, 'smirnof', 0, 'duapuluh', 'limaribu'),
(23, 'bintang', 0, 'lima tiga', 'duapuluhjuta'),
(24, 'anggur dong', 0, 'waduh', 'delapan'),
(25, 'morgan', 0, 'sdsd', 'awdffaw');

-- --------------------------------------------------------

--
-- Struktur dari tabel `categories`
--

CREATE TABLE `categories` (
  `category_id` bigint(20) UNSIGNED NOT NULL,
  `category_name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data untuk tabel `categories`
--

INSERT INTO `categories` (`category_id`, `category_name`, `description`, `created_at`, `updated_at`) VALUES
(1, 'cok', 'njir', NULL, NULL),
(2, 'Condiments', 'Sweet and savory sauces, relishes, spreads, and seasonings', NULL, NULL),
(3, 'Confections', 'Desserts, candies, and sweet breads', NULL, NULL),
(4, 'Dairy Products', 'Cheeses', NULL, NULL),
(5, 'Grains/Cereals', 'Breads, crackers, pasta, and cereal', NULL, NULL),
(6, 'Meat/Poultry', 'Prepared meats', NULL, NULL),
(7, 'Produce', 'Dried fruit and bean curd', NULL, NULL),
(8, 'Seafood', 'Seaweed and fish', NULL, NULL),
(9, 'ayam', 'wdwdwd', NULL, NULL),
(10, 'ayam', 'kuda', NULL, NULL),
(12, 'Beverages', 'Soft drinks, coffees, teas, beers, and ales', NULL, NULL),
(13, 'Beverages', 'Soft drinks, coffees, teas, beers, and ales', NULL, NULL),
(18, 'egegegeg', 'egegeeww', NULL, NULL),
(19, 'ayam', 'sapi', NULL, NULL),
(20, 'bir', 'kane nih bergizi', NULL, NULL),
(21, 'enak enak', 'kane nih', NULL, NULL),
(22, 'enak enak', 'kane nih', NULL, NULL),
(23, 'xxx', 'puyeng', NULL, NULL),
(24, 'kane nih', 'joss', NULL, NULL),
(25, 'awdasd', 'ffefef', NULL, NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `jenis`
--

CREATE TABLE `jenis` (
  `id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `products`
--

CREATE TABLE `products` (
  `id` bigint(20) NOT NULL,
  `nama_product` varchar(255) DEFAULT NULL,
  `deskripsi` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `products`
--

INSERT INTO `products` (`id`, `nama_product`, `deskripsi`) VALUES
(1, 'motor', 'sepeda motor astrea'),
(2, 'hiu balap', 'sepeda dari pak prabwo dan jokowi');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `apikey` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `apikey`) VALUES
(1, 'septian', 'septian@gmail.com', 'nugraha', '7535158ce69cbb59ac30cde46a8dff6f7158f722204675e1a8e12bf7b8d90e35'),
(2, 'nugraha', 'nugraha@gmail.com', 'nugraha', '88bebbfceda03451abf1d337778c53250ea24b65d7d3e8b9e0b2b808230c624c'),
(3, 'nugraha', 'nugraha@gmail.com', 'nugraha', '81662f22a63bc49f5e1215be5b336e67f7a312f749b1b0c9efe7369144133313'),
(4, 'nugraha', 'nugraha@gmail.com', 'nugraha', '6441125fef6b411d649edcffee5114330b008e23becb527b077d9d2c0422012f'),
(5, 'septian21', 'septiann@gmail.com', 'nugraha', 'e761534f241f8ae378c8a6598cf2700763f5cc3c6290e83c8ef40f9976a63167');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `barangs`
--
ALTER TABLE `barangs`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`category_id`);

--
-- Indeks untuk tabel `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `barangs`
--
ALTER TABLE `barangs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT untuk tabel `categories`
--
ALTER TABLE `categories`
  MODIFY `category_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT untuk tabel `products`
--
ALTER TABLE `products`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
