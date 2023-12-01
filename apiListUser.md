## Intro

| Method                                                           | Url               | Description         |
| ---------------------------------------------------------------- | ----------------- | ------------------- |
| ![](https://storage.kodeteks.com/POST.png)                       | `/register`       | Register Users      |
| ![](https://storage.kodeteks.com/POST.png)                       | `/login`          | Login Users         |
| ![](https://storage.kodeteks.com/POST.png)                       | `/forgot-password`| Forgot Password     |
| ![](https://storage.kodeteks.com/POST.png)                       | `/verify-otp`     | Forgot Password     |
| ![](https://storage.kodeteks.com/POST.png)                       | `/new-password`   | Forgot Password     |

## Home & Info

| Method                                                           | Url                      | Description                                        |
| ---------------------------------------------------------------- | ------------------------ | -------------------------------------------------- |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/home`                  | View All Data (point,article popular,mission event and achivement)|
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/faq`                   | View All FAQ            |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/faq/:id`               | View Detail FAQ         |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/achievements`          | View All Achievement    |
| ![](https://storage.kodeteks.com/POST.png)                       | `/recybot`               | Chatbot RecyThing       |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/notifications`         | View User Notifications |

## User Profile

| Method                                                           | Url                            | Description             |
| ---------------------------------------------------------------- | ------------------------------ | ------------------------|
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/users/profile`               | View User Profile       |
| ![](https://storage.kodeteks.com/PUT.png)                        | `/users/profile`               | Edit User Profile       |
| ![](https://storage.kodeteks.com/PATCH.png)                      | `/users/profile/reset-password`| Reset Password          |

## Users Point & Voucher

| Method                                                           | Url                             | Description                            |
| ---------------------------------------------------------------- | ------------------------------- | -------------------------------------- |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/users/point`                  | View User Point                        |
| ![](https://storage.kodeteks.com/POST.png)                       | `/users/point/daily`            | Claim Daily Point                      |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/users/point/history`          | View All History Point (tukar point, claim point, mission point)                |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/users/point/history/:id`      | View Detail History Point              |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/vouchers`               | View All Voucher                       |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/vouchers/:id`           | View Detail Voucher                    |
| ![](https://storage.kodeteks.com/POST.png)                       | `/vouchers/:id`           | Confirm Exchange Point With Email      |

## Article & Drop Point

### Article
| Method                                                           | Url                     | Query Param          | Description                    |
| ---------------------------------------------------------------- | ----------------------- | -------------------- | ------------------------------ |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/articles`             | `category`,`search`  | View Articles Popular,category,search          |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/articles/:id`         |                      | View Detail Articles           |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/articles/categories`  |                      | View All Categories Article    |

### Drop Point
| Method                                                           | Url                     | Query Param    | Description                    |
| ---------------------------------------------------------------- | ----------------------- | -------------- | ------------------------------ |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/drop-point`           | `search`       | View All Drop Point and Search |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/drop-point/:id`       |                | View Detail Articles           |

## Users Report
| Method                                                           | Url                         | Description                            |
| ---------------------------------------------------------------- | --------------------------- | -------------------------------------- |
| ![](https://storage.kodeteks.com/POST.png)                       | `/users/report`             | View User Point                        |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/users/report/history`     | View User History                      |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/users/report/history/:id` | View Detail History Point              |

## Community
| Method                                                           | Url                         | Query Param               | Description                            |
| ---------------------------------------------------------------- | --------------------------- | ------------------------- | -------------------------------------- |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/comunities`               | `location`,`most_members`,`search` | View User Comunity and recomendation   |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/comunities/:id`           |                           | View Detail Community,only member can see event   |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/comunities/users`         |                           | View Community Users Following |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/comunities/recommendation`|                           | View Community Recomendation   |
| ![](https://storage.kodeteks.com/POST.png)                       | `/comunities/:id`           |                           | Follow Community               |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/comunities/event/:id`     |                           | View Community Event, only member can see   |

### Mission 
| Method                                                           | Url                     | Query Param    | Description                    |
| ---------------------------------------------------------------- | ----------------------- | -------------- | ------------------------------ |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/missions`             | `status`       | View All Available Missions,  |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/missions/:id`       |                | View Detail mission           |
| ![](https://storage.kodeteks.com/POST.png)  | `/missions`       |                | Accept Challenge           |
| ![](https://storage.kodeteks.com/POST.png)  | `/missions/proof`       |                | upload proof           |

