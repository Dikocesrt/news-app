# TEST JUNIOR BACKEND INDONESIA CAKAP DIGITAL NEWS API

<div align="center">
  <a href="https://www.instagram.com/hightechteacher_id/">
    <img src="https://perusahaan.net/foto/1250566/logo-pt-indonesia-cakap-digital.jpg" alt="Logo">
  </a>
</div>

> This project is a backend developer test project using Clean Architecture Golang Echo at PT Indonesia Cakap Digital.

# About Project

This project is a CMS (Content Management System) backend application that provides features such as user authentication, news management with categories, custom page creation, and anonymous commenting. The system distinguishes between authenticated users and guests, and delivers output in the form of a RESTful API. Authenticated users can manage news, categories, and pages, while guests can view news content and submit anonymous comments.

# Target Scope

-   🧍 **User**: Can log in, create and manage news, categories, and custom pages.
-   👤 **Guest**: Can view news and submit anonymous comments.

# Features:

### List Features:

| Method | Feature                 |
| ------ | ----------------------- |
| POST   | Register user           |
| POST   | Login user              |
| POST   | Create Category         |
| GET    | Fetch All Categories    |
| GET    | Fetch Category By ID    |
| PUT    | Update Category         |
| DELETE | Delete Category         |
| POST   | Create News             |
| GET    | Fetch All News          |
| GET    | Fetch News By ID        |
| PUT    | Update News             |
| DELETE | Delete News             |
| POST   | Create Comment          |
| POST   | Create Custom Page      |
| GET    | Fetch All Custom Pages  |
| GET    | Fetch Custom Page By ID |
| PUT    | Update Custom Page      |
| DELETE | Delete Custom Page      |

# TECH STACK

-   [Golang](https://github.com/golang/go) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software!
-   [Echo](https://github.com/labstack/echo) - High performance, extensible, minimalist Go web framework.
-   [GORM](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang, aims to be developer friendly.
-   [MySQL](https://www.mysql.com/) - For Database

# Getting Started

### Git Clone

```sh
git clone https://github.com/Dikocesrt/news-app.git
```

### Create an `.env` file locally. Or You can duplicate `.env example`

```sh
cp .env.example .env
```

### Install required dependencies

This command ensures that the depedency matches the source code in your module's directory, adding any missing dependencies and removing unnecessary ones.

```sh
go mod tidy
```

### Run the app

Execute the following command to run the main.go file:

```sh
go run main.go
```

# API Documentation

This is the API documentation for Test Indonesia Cakap Digital News App. This document provides an overview of the endpoints, request methods, parameters, and responses supported by our API.

_For more examples, please refer to the [Documentation](https://documenter.getpostman.com/view/27063468/2sB2j6AqUV)_

# Entity Relationship Diagram

<div align="center">
  <a href="https://res.cloudinary.com/dy2fwknbn/image/upload/v1746499749/Screenshot_2025-05-06_at_09.47.47_j81n5j.png">
    <img src="https://res.cloudinary.com/dy2fwknbn/image/upload/v1746499749/Screenshot_2025-05-06_at_09.47.47_j81n5j.png" alt="Logo">
  </a>
</div>

# Contributors

**MAHARDIKO CESARTISTA RASENDRIYA**
<br>
[![MAHARDIKO CESARTISTA RASENDRIYA - GitHub](https://img.shields.io/badge/MAHARDIKO_CESARTISTA_RASENDRIYA-black?logo=github)](https://github.com/Dikocesrt)
