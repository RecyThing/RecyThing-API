## Login Admin

| Method                                                           | Url                   | Description                        |
| ---------------------------------------------------------------- | --------------------- | ---------------------------------- |
| ![](https://storage.kodeteks.com/POST.png)                       | `/admins/login`       | Login Admin & Super admin          |


## Manage Admin (super admin)

| Method                                                           | Url             | Description         |
| ---------------------------------------------------------------- | --------------- | ------------------- |
| ![](https://storage.kodeteks.com/POST.png)                       | `/admins`       | Add Admin           |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/admins`   | View All Admin      |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png) | `/admins/:id`   | View Admin By Id    |
| ![](https://storage.kodeteks.com/PUT.png)                        | `/admins/:id`   | Update Admin By Id  |
| ![](https://storage.kodeteks.com/DELETE.png)                     | `/admins/:id`   | Delete Admin By Id  |


## Manage User

| Method                                                              | Url                          | Description         |
| ------------------------------------------------------------------- | ---------------------------- | ------------------- |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/users`       | View All Users      |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/users/:id`   | View User By Id     |
| ![](https://storage.kodeteks.com/DELETE.png)                        | `/admins/manage/users/:id`   | Delete User By Id   |

## Dashboard

| Method                                                              | Url                            | Description                          |
| ------------------------------------------------------------------- | ------------------------------ | ------------------------------------ |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/dashboard`            | View All Dashboard                   |

### Dashboard Report Download

| Method                                                              | Url                                | Description                          |
| ------------------------------------------------------------------- | ---------------------------------- | ------------------------------------ |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/dashboard/download`       | View All Report Download             |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/dashboard/download/:file` | Get Report Download                  |

## Manage Reward

| Method                                                              | Url                            | Description                          |
| ------------------------------------------------------------------- | ------------------------------ | ------------------------------------ |
| ![](https://storage.kodeteks.com/POST.png)                          | `/admins/manage/rewards`       | Add Reward                           |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/rewards`       | View All Reward                      |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/rewards/:id`   | View Detail Reward                   |
| ![](https://storage.kodeteks.com/PUT.png)                           | `/admins/manage/rewards/:id`   | Update Reward By Id                  |
| ![](https://storage.kodeteks.com/DELETE.png)                        | `/admins/manage/rewards/:id`   | Delete Reward By Id                  |

## Manage Report

| Method                                                              | Url                            | Query Param | Description                          |
| ------------------------------------------------------------------- | ------------------------------ | ----------- | ------------------------------------ |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/reports`       | `status`    | View All Report & filter status      |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/reports/:id`   |             | View Detail Report                   |
| ![](https://storage.kodeteks.com/PATCH.png)                         | `/admins/manage/reports/:id`   |             | Update Report                        |

## Manage Achievement

| Method                                                              | Url                                | Description                          |
| ------------------------------------------------------------------- | ---------------------------------- | ------------------------------------ |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/achivements`       | View All Achievements                |
| ![](https://storage.kodeteks.com/PATCH.png)                         | `/admins/manage/achivements/:id`   | Update Achievement                   |

## Manage Articles

| Method                                                              | Url                             | Description                          |
| ------------------------------------------------------------------- | ------------------------------  | ------------------------------------ |
| ![](https://storage.kodeteks.com/POST.png)                          | `/admins/manage/articles`       | Add Content                          |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/articles`       | View All articles                    |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/articles/:id`   | View Detail Content                  |
| ![](https://storage.kodeteks.com/PUT.png)                           | `/admins/manage/articles/:id`   | Update articles By Id                 |
| ![](https://storage.kodeteks.com/DELETE.png)                        | `/admins/manage/articles/:id`   | Delete articles By Id                 |

## Manage Mission

| Method                                                              | Url                            | Query Param | Description            |
| ------------------------------------------------------------------- | ------------------------------ | ----------- | ---------------------- |
| ![](https://storage.kodeteks.com/POST.png)                          | `/admins/manage/missions`      |             | Add Mission            |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/missions`      | `status`    | View All Missions      |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/missions/:id`  |             | View Detail Mission    |
| ![](https://storage.kodeteks.com/PUT.png)                           | `/admins/manage/missions/:id`  |             | Update Mission By Id   |
| ![](https://storage.kodeteks.com/DELETE.png)                        | `/admins/manage/missions/:id`  |             | Delete Mission By Id   |

### Mission Approval

| Method                                                              | Url                                      | Query Param | Description             |
| ------------------------------------------------------------------- | ---------------------------------------- | ----------- | ----------------------- |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/missions/approval`       | `status`    | View All Contents       |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/missions/approval/:id`   |             | View Detail Content     |
| ![](https://storage.kodeteks.com/PATCH.png)                         | `/admins/manage/missions/approval/:id`   |             | Update Content By Id    |

## Manage Prompt

| Method                                                              | Url                            | Query Param | Description                          |
| ------------------------------------------------------------------- | ------------------------------ | ----------- | ------------------------------------ |
| ![](https://storage.kodeteks.com/POST.png)                          | `/admins/manage/prompts`       |             | Add Prompt                           |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/prompts`       | `topic`     | View All Prompts & filter topic      |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/prompts/:id`   |             | View Detail Prompt                   |
| ![](https://storage.kodeteks.com/PUT.png)                           | `/admins/manage/prompts/:id`   |             | Update Prompt By Id                  |
| ![](https://storage.kodeteks.com/DELETE.png)                        | `/admins/manage/prompts/:id`   |             | Delete Prompt By Id                  |

## Manage Drop Point

| Method                                                              | Url                                | Description                    |
| ------------------------------------------------------------------- | ---------------------------------- | ------------------------------ |
| ![](https://storage.kodeteks.com/POST.png)                          | `/admins/manage/drop-points`       | Add Drop Point                 |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/drop-points`       | View All Drop Points           |
| ![](https://storage.kodeteks.com/PUT.png)                           | `/admins/manage/drop-points/:id`   | Update Drop Point By Id        |
| ![](https://storage.kodeteks.com/DELETE.png)                        | `/admins/manage/drop-points/:id`   | Delete Drop Point By Id        |

## Manage Penukaran Sampah

| Method                                                              | Url                                | Description                    |
| ------------------------------------------------------------------- | ---------------------------------- | ------------------------------ |
| ![](https://storage.kodeteks.com/POST.png)                          | `/admins/manage/recycles`          | Add Recycle                    |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/recycles`          | View All Recycles              |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/recycles/:id`      | View Detail Recycle            |
| ![](https://storage.kodeteks.com/DELETE.png)                        | `/admins/manage/recycles/:id`      | Delete Recycle By Id           |

## Manage Trash Category

| Method                                                              | Url                                | Description            |
| ------------------------------------------------------------------- | ---------------------------------- | ---------------------- |
| ![](https://storage.kodeteks.com/POST.png)                          | `/admins/manage/trashes`           | Add Trash              |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/trashes`           | View All Trash         |
| ![](https://storage.kodeteks.com/PUT.png)                           | `/admins/manage/trashes/:id`       | Update Trash By Id     |
| ![](https://storage.kodeteks.com/DELETE.png)                        | `/admins/manage/trashes/:id`       | Delete Trash By Id     |

## Manage Exchange Point

| Method                                                              | Url                                    | Query Param | Description                          |
| ------------------------------------------------------------------- | -------------------------------------- | ----------- | ------------------------------------ |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/exchange-points`       | `status`    | View All Exchange Point & filter     |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/exchange-points/:id`   |             | View Detail Exchange Point           |
| ![](https://storage.kodeteks.com/PATCH.png)                         | `/admins/manage/exchange-points/:id`   |             | Update Exchange Point By Id          |

## Manage Community

| Method                                                              | Url                                | Description                    |
| ------------------------------------------------------------------- | ---------------------------------- | ------------------------------ |
| ![](https://storage.kodeteks.com/POST.png)                          | `/admins/manage/communities`       | Add Community                  |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/communities`       | View All Communities           |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/communities/:id`   | View Detail Community          |
| ![](https://storage.kodeteks.com/PUT.png)                           | `/admins/manage/communities/:id`   | Update Community By Id         |
| ![](https://storage.kodeteks.com/DELETE.png)                        | `/admins/manage/communities/:id`   | Delete Community By Id         |

### Community Event

| Method                                                              | Url                                      | Query Param | Description                    |
| ------------------------------------------------------------------- | ---------------------------------------- | ----------- | ------------------------------ |
| ![](https://storage.kodeteks.com/POST.png)                          | `/admins/manage/communities/event`       |             | Add Community Event            |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/communities/event`       | `status`    | View All Community Events      |
| ![](https://pub-cc8247a7807d42d1bd2453b3dae2f678.r2.dev/GET.png)    | `/admins/manage/communities/event/:id`   |             | View Detail Community Event    |
| ![](https://storage.kodeteks.com/PUT.png)                           | `/admins/manage/communities/event/:id`   |             | Update Community Event By Id   |
| ![](https://storage.kodeteks.com/DELETE.png)                        | `/admins/manage/communities/event/:id`   |             | Delete Community Event By Id   |


