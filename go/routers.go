/*
 * Документация публичного API
 *
 * # Введение API Timeweb Cloud позволяет вам управлять ресурсами в облаке программным способом с использованием обычных HTTP-запросов.  Множество функций, которые доступны в панели управления Timeweb Cloud, также доступны через API, что позволяет вам автоматизировать ваши собственные сценарии.  В этой документации сперва будет описан общий дизайн и принципы работы API, а после этого конкретные конечные точки. Также будут приведены примеры запросов к ним.   ## Запросы Запросы должны выполняться по протоколу `HTTPS`, чтобы гарантировать шифрование транзакций. Поддерживаются следующие методы запроса: |Метод|Применение| |--- |--- | |GET|Извлекает данные о коллекциях и отдельных ресурсах.| |POST|Для коллекций создает новый ресурс этого типа. Также используется для выполнения действий с конкретным ресурсом.| |PUT|Обновляет существующий ресурс.| |PATCH|Некоторые ресурсы поддерживают частичное обновление, то есть обновление только части атрибутов ресурса, в этом случае вместо метода PUT будет использован PATCH.| |DELETE|Удаляет ресурс.|  Методы `POST`, `PUT` и `PATCH` могут включать объект в тело запроса с типом содержимого `application/json`.  ### Параметры в запросах Некоторые коллекции поддерживают пагинацию, поиск или сортировку в запросах. В параметрах запроса требуется передать: - `limit` — обозначает количество записей, которое необходимо вернуть  - `offset` — указывает на смещение, относительно начала списка  - `search` — позволяет указать набор символов для поиска  - `sort` — можно задать правило сортировки коллекции  ## Ответы Запросы вернут один из следующих кодов состояния ответа HTTP:  |Статус|Описание| |--- |--- | |200 OK|Действие с ресурсом было выполнено успешно.| |201 Created|Ресурс был успешно создан. При этом ресурс может быть как уже готовым к использованию, так и находиться в процессе запуска.| |204 No Content|Действие с ресурсом было выполнено успешно, и ответ не содержит дополнительной информации в теле.| |400 Bad Request|Был отправлен неверный запрос, например, в нем отсутствуют обязательные параметры и т. д. Тело ответа будет содержать дополнительную информацию об ошибке.| |401 Unauthorized|Ошибка аутентификации.| |403 Forbidden|Аутентификация прошла успешно, но недостаточно прав для выполнения действия.| |404 Not Found|Запрашиваемый ресурс не найден.| |409 Conflict|Запрос конфликтует с текущим состоянием.| |423 Locked|Ресурс из запроса заблокирован от применения к нему указанного метода.| |429 Too Many Requests|Был достигнут лимит по количеству запросов в единицу времени.| |500 Internal Server Error|При выполнении запроса произошла какая-то внутренняя ошибка. Чтобы решить эту проблему, лучше всего создать тикет в панели управления.|  ### Структура успешного ответа Все конечные точки будут возвращать данные в формате `JSON`. Ответы на `GET`-запросы будут иметь на верхнем уровне следующую структуру атрибутов:  |Название поля|Тип|Описание| |--- |--- |--- | |[entity_name]|object, object[], string[], number[], boolean|Динамическое поле, которое будет меняться в зависимости от запрашиваемого ресурса и будет содержать все атрибуты, необходимые для описания этого ресурса. Например, при запросе списка баз данных будет возвращаться поле `dbs`, а при запросе конкретного облачного сервера `server`. Для некоторых конечных точек в ответе может возвращаться сразу несколько ресурсов.| |meta|object|Опционально. Объект, который содержит вспомогательную информацию о ресурсе. Чаще всего будет встречаться при запросе коллекций и содержать поле `total`, которое будет указывать на количество элементов в коллекции.| |response_id|string|Опционально. В большинстве случаев в ответе будет содержаться уникальный идентификатор ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот идентификатор — так мы сможем найти ответ на него намного быстрее. Также вы можете использовать этот идентификатор, чтобы убедиться, что это новый ответ на запрос и результат не был получен из кэша.|  Пример запроса на получение списка SSH-ключей: ```     HTTP/2.0 200 OK     {       \"ssh_keys\":[           {             \"body\":\"ssh-rsa AAAAB3NzaC1sdfghjkOAsBwWhs= example@device.local\",             \"created_at\":\"2021-09-15T19:52:27Z\",             \"expired_at\":null,             \"id\":5297,             \"is_default\":false,             \"name\":\"example@device.local\",             \"used_at\":null,             \"used_by\":[]           }       ],       \"meta\":{           \"total\":1       },       \"response_id\":\"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"     } ```  ### Структура ответа с ошибкой |Название поля|Тип|Описание| |--- |--- |--- | |status_code|number|Короткий числовой идентификатор ошибки.| |error_code|string|Короткий текстовый идентификатор ошибки, который уточняет числовой идентификатор и удобен для программной обработки. Самый простой пример — это код `not_found` для ошибки 404.| |message|string, string[]|Опционально. В большинстве случаев в ответе будет содержаться человекочитаемое подробное описание ошибки или ошибок, которые помогут понять, что нужно исправить.| |response_id|string|Опционально. В большинстве случае в ответе будет содержаться уникальный идентификатор ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот идентификатор — так мы сможем найти ответ на него намного быстрее.|  Пример: ```     HTTP/2.0 403 Forbidden     {       \"status_code\": 403,       \"error_code\":  \"forbidden\",       \"message\":     \"You do not have access for the attempted action\",       \"response_id\": \"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"     } ```  ## Статусы ресурсов Важно учесть, что при создании большинства ресурсов внутри платформы вам будет сразу возвращен ответ от сервера со статусом `200 OK` или `201 Created` и идентификатором созданного ресурса в теле ответа, но при этом этот ресурс может быть ещё в *состоянии запуска*.  Для того чтобы понять, в каком состоянии сейчас находится ваш ресурс, мы добавили поле `status` в ответ на получение информации о ресурсе.  Список статусов будет отличаться в зависимости от типа ресурса. Увидеть поддерживаемый список статусов вы сможете в описании каждого конкретного ресурса.     ## Ограничение скорости запросов (Rate Limiting) Чтобы обеспечить стабильность для всех пользователей, Timeweb Cloud защищает API от всплесков входящего трафика, анализируя количество запросов c каждого аккаунта к каждой конечной точке.  Если ваше приложение отправляет более 20 запросов в секунду на одну конечную точку, то для этого запроса API может вернуть код состояния HTTP `429 Too Many Requests`.   ## Аутентификация Доступ к API осуществляется с помощью JWT-токена. Токенами можно управлять внутри панели управления Timeweb Cloud в разделе *API и Terraform*.  Токен необходимо передавать в заголовке каждого запроса в формате: ```   Authorization: Bearer $TIMEWEB_CLOUD_TOKEN ```  ## Формат примеров API Примеры в этой документации описаны с помощью `curl`, HTTP-клиента командной строки. На компьютерах `Linux` и `macOS` обычно по умолчанию установлен `curl`, и он доступен для загрузки на всех популярных платформах, включая `Windows`.  Каждый пример разделен на несколько строк символом `\\`, который совместим с `bash`. Типичный пример выглядит так: ```   curl -X PATCH      -H \"Content-Type: application/json\"      -H \"Authorization: Bearer $TIMEWEB_CLOUD_TOKEN\"      -d '{\"name\":\"Cute Corvus\",\"comment\":\"Development Server\"}'      \"https://api.timeweb.cloud/api/v1/dedicated/1051\" ``` - Параметр `-X` задает метод запроса. Для согласованности метод будет указан во всех примерах, даже если он явно не требуется для методов `GET`. - Строки `-H` задают требуемые HTTP-заголовки. - Примеры, для которых требуется объект JSON в теле запроса, передают требуемые данные через параметр `-d`.  Чтобы использовать приведенные примеры, не подставляя каждый раз в них свой токен, вы можете добавить токен один раз в переменные окружения в вашей консоли. Например, на `Linux` это можно сделать с помощью команды:  ``` TIMEWEB_CLOUD_TOKEN=\"token\" ```  После этого токен будет автоматически подставляться в ваши запросы.  Обратите внимание, что все значения в этой документации являются примерами. Не полагайтесь на идентификаторы операционных систем, тарифов и т.д., используемые в примерах. Используйте соответствующую конечную точку для получения значений перед созданием ресурсов.   ## Версионирование API построено согласно принципам [семантического версионирования](https://semver.org/lang/ru). Это значит, что мы гарантируем обратную совместимость всех изменений в пределах одной мажорной версии.  Мажорная версия каждой конечной точки обозначается в пути запроса, например, запрос `/api/v1/servers` указывает, что этот метод имеет версию 1.
 *
 * API version: 1.0.0
 * Contact: info@timeweb.cloud
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AddBalancerToProject",
		strings.ToUpper("Post"),
		"/api/v1/projects/{project_id}/resources/balancers",
		AddBalancerToProject,
	},

	Route{
		"AddClusterToProject",
		strings.ToUpper("Post"),
		"/api/v1/projects/{project_id}/resources/clusters",
		AddClusterToProject,
	},

	Route{
		"AddCountriesToAllowedList",
		strings.ToUpper("Post"),
		"/api/v1/auth/access/countries",
		AddCountriesToAllowedList,
	},

	Route{
		"AddDatabaseToProject",
		strings.ToUpper("Post"),
		"/api/v1/projects/{project_id}/resources/databases",
		AddDatabaseToProject,
	},

	Route{
		"AddDedicatedServerToProject",
		strings.ToUpper("Post"),
		"/api/v1/projects/{project_id}/resources/dedicated",
		AddDedicatedServerToProject,
	},

	Route{
		"AddDomain",
		strings.ToUpper("Post"),
		"/api/v1/add-domain/{fqdn}",
		AddDomain,
	},

	Route{
		"AddIPsToAllowedList",
		strings.ToUpper("Post"),
		"/api/v1/auth/access/ips",
		AddIPsToAllowedList,
	},

	Route{
		"AddIPsToBalancer",
		strings.ToUpper("Post"),
		"/api/v1/balancers/{balancer_id}/ips",
		AddIPsToBalancer,
	},

	Route{
		"AddServerToProject",
		strings.ToUpper("Post"),
		"/api/v1/projects/{project_id}/resources/servers",
		AddServerToProject,
	},

	Route{
		"AddStorageToProject",
		strings.ToUpper("Post"),
		"/api/v1/projects/{project_id}/resources/buckets",
		AddStorageToProject,
	},

	Route{
		"AddSubdomain",
		strings.ToUpper("Post"),
		"/api/v1/domains/{fqdn}/subdomains/{subdomain_fqdn}",
		AddSubdomain,
	},

	Route{
		"CheckDomain",
		strings.ToUpper("Get"),
		"/api/v1/check-domain/{fqdn}",
		CheckDomain,
	},

	Route{
		"CreateBalancer",
		strings.ToUpper("Post"),
		"/api/v1/balancers",
		CreateBalancer,
	},

	Route{
		"CreateBalancerRule",
		strings.ToUpper("Post"),
		"/api/v1/balancers/{balancer_id}/rules",
		CreateBalancerRule,
	},

	Route{
		"CreateDomainDNSRecord",
		strings.ToUpper("Post"),
		"/api/v1/domains/{fqdn}/dns-records",
		CreateDomainDNSRecord,
	},

	Route{
		"CreateDomainMailbox",
		strings.ToUpper("Post"),
		"/api/v1/mail/domains/{domain}",
		CreateDomainMailbox,
	},

	Route{
		"CreateDomainRequest",
		strings.ToUpper("Post"),
		"/api/v1/domains-requests",
		CreateDomainRequest,
	},

	Route{
		"CreateImage",
		strings.ToUpper("Post"),
		"/api/v1/images",
		CreateImage,
	},

	Route{
		"CreateImageDownloadUrl",
		strings.ToUpper("Post"),
		"/api/v1/images/{image_id}/download-url",
		CreateImageDownloadUrl,
	},

	Route{
		"CreateProject",
		strings.ToUpper("Post"),
		"/api/v1/projects",
		CreateProject,
	},

	Route{
		"DeleteBalancer",
		strings.ToUpper("Delete"),
		"/api/v1/balancers/{balancer_id}",
		DeleteBalancer,
	},

	Route{
		"DeleteBalancerRule",
		strings.ToUpper("Delete"),
		"/api/v1/balancers/{balancer_id}/rules/{rule_id}",
		DeleteBalancerRule,
	},

	Route{
		"DeleteCountriesFromAllowedList",
		strings.ToUpper("Delete"),
		"/api/v1/auth/access/countries",
		DeleteCountriesFromAllowedList,
	},

	Route{
		"DeleteDomain",
		strings.ToUpper("Delete"),
		"/api/v1/domains/{fqdn}",
		DeleteDomain,
	},

	Route{
		"DeleteDomainDNSRecord",
		strings.ToUpper("Delete"),
		"/api/v1/domains/{fqdn}/dns-records/{record_id}",
		DeleteDomainDNSRecord,
	},

	Route{
		"DeleteIPsFromAllowedList",
		strings.ToUpper("Delete"),
		"/api/v1/auth/access/ips",
		DeleteIPsFromAllowedList,
	},

	Route{
		"DeleteIPsFromBalancer",
		strings.ToUpper("Delete"),
		"/api/v1/balancers/{balancer_id}/ips",
		DeleteIPsFromBalancer,
	},

	Route{
		"DeleteImage",
		strings.ToUpper("Delete"),
		"/api/v1/images/{image_id}",
		DeleteImage,
	},

	Route{
		"DeleteImageDownloadURL",
		strings.ToUpper("Delete"),
		"/api/v1/images/{image_id}/download-url/{image_url_id}",
		DeleteImageDownloadURL,
	},

	Route{
		"DeleteMailbox",
		strings.ToUpper("Delete"),
		"/api/v1/mail/domains/{domain}/mailboxes/{mailbox}",
		DeleteMailbox,
	},

	Route{
		"DeleteProject",
		strings.ToUpper("Delete"),
		"/api/v1/projects/{project_id}",
		DeleteProject,
	},

	Route{
		"DeleteSubdomain",
		strings.ToUpper("Delete"),
		"/api/v1/domains/{fqdn}/subdomains/{subdomain_fqdn}",
		DeleteSubdomain,
	},

	Route{
		"GetAccountBalancers",
		strings.ToUpper("Get"),
		"/api/v1/projects/resources/balancers",
		GetAccountBalancers,
	},

	Route{
		"GetAccountClusters",
		strings.ToUpper("Get"),
		"/api/v1/projects/resources/clusters",
		GetAccountClusters,
	},

	Route{
		"GetAccountDatabases",
		strings.ToUpper("Get"),
		"/api/v1/projects/resources/databases",
		GetAccountDatabases,
	},

	Route{
		"GetAccountDedicatedServers",
		strings.ToUpper("Get"),
		"/api/v1/projects/resources/dedicated",
		GetAccountDedicatedServers,
	},

	Route{
		"GetAccountServers",
		strings.ToUpper("Get"),
		"/api/v1/projects/resources/servers",
		GetAccountServers,
	},

	Route{
		"GetAccountStatus",
		strings.ToUpper("Get"),
		"/api/v1/account/status",
		GetAccountStatus,
	},

	Route{
		"GetAccountStorages",
		strings.ToUpper("Get"),
		"/api/v1/projects/resources/buckets",
		GetAccountStorages,
	},

	Route{
		"GetAllProjectResources",
		strings.ToUpper("Get"),
		"/api/v1/projects/{project_id}/resources",
		GetAllProjectResources,
	},

	Route{
		"GetAuthAccessSettings",
		strings.ToUpper("Get"),
		"/api/v1/auth/access",
		GetAuthAccessSettings,
	},

	Route{
		"GetBalancer",
		strings.ToUpper("Get"),
		"/api/v1/balancers/{balancer_id}",
		GetBalancer,
	},

	Route{
		"GetBalancerIPs",
		strings.ToUpper("Get"),
		"/api/v1/balancers/{balancer_id}/ips",
		GetBalancerIPs,
	},

	Route{
		"GetBalancerRules",
		strings.ToUpper("Get"),
		"/api/v1/balancers/{balancer_id}/rules",
		GetBalancerRules,
	},

	Route{
		"GetBalancers",
		strings.ToUpper("Get"),
		"/api/v1/balancers",
		GetBalancers,
	},

	Route{
		"GetBalancersPresets",
		strings.ToUpper("Get"),
		"/api/v1/presets/balancers",
		GetBalancersPresets,
	},

	Route{
		"GetCountries",
		strings.ToUpper("Get"),
		"/api/v1/auth/access/countries",
		GetCountries,
	},

	Route{
		"GetDomain",
		strings.ToUpper("Get"),
		"/api/v1/domains/{fqdn}",
		GetDomain,
	},

	Route{
		"GetDomainDNSRecords",
		strings.ToUpper("Get"),
		"/api/v1/domains/{fqdn}/dns-records",
		GetDomainDNSRecords,
	},

	Route{
		"GetDomainDefaultDNSRecords",
		strings.ToUpper("Get"),
		"/api/v1/domains/{fqdn}/default-dns-records",
		GetDomainDefaultDNSRecords,
	},

	Route{
		"GetDomainMailInfo",
		strings.ToUpper("Get"),
		"/api/v1/mail/domains/{domain}/info",
		GetDomainMailInfo,
	},

	Route{
		"GetDomainMailboxes",
		strings.ToUpper("Get"),
		"/api/v1/mail/domains/{domain}",
		GetDomainMailboxes,
	},

	Route{
		"GetDomainNameServers",
		strings.ToUpper("Get"),
		"/api/v1/domains/{fqdn}/name-servers",
		GetDomainNameServers,
	},

	Route{
		"GetDomainRequest",
		strings.ToUpper("Get"),
		"/api/v1/domains-requests/{request_id}",
		GetDomainRequest,
	},

	Route{
		"GetDomainRequests",
		strings.ToUpper("Get"),
		"/api/v1/domains-requests",
		GetDomainRequests,
	},

	Route{
		"GetDomains",
		strings.ToUpper("Get"),
		"/api/v1/domains",
		GetDomains,
	},

	Route{
		"GetFinances",
		strings.ToUpper("Get"),
		"/api/v1/account/finances",
		GetFinances,
	},

	Route{
		"GetImage",
		strings.ToUpper("Get"),
		"/api/v1/images/{image_id}",
		GetImage,
	},

	Route{
		"GetImageDownloadURL",
		strings.ToUpper("Get"),
		"/api/v1/images/{image_id}/download-url/{image_url_id}",
		GetImageDownloadURL,
	},

	Route{
		"GetImageDownloadURLs",
		strings.ToUpper("Get"),
		"/api/v1/images/{image_id}/download-url",
		GetImageDownloadURLs,
	},

	Route{
		"GetImages",
		strings.ToUpper("Get"),
		"/api/v1/images",
		GetImages,
	},

	Route{
		"GetMailQuota",
		strings.ToUpper("Get"),
		"/api/v1/mail/quota",
		GetMailQuota,
	},

	Route{
		"GetMailbox",
		strings.ToUpper("Get"),
		"/api/v1/mail/domains/{domain}/mailboxes/{mailbox}",
		GetMailbox,
	},

	Route{
		"GetMailboxes",
		strings.ToUpper("Get"),
		"/api/v1/mail",
		GetMailboxes,
	},

	Route{
		"GetNotificationSettings",
		strings.ToUpper("Get"),
		"/api/v1/account/notification-settings",
		GetNotificationSettings,
	},

	Route{
		"GetProject",
		strings.ToUpper("Get"),
		"/api/v1/projects/{project_id}",
		GetProject,
	},

	Route{
		"GetProjectBalancers",
		strings.ToUpper("Get"),
		"/api/v1/projects/{project_id}/resources/balancers",
		GetProjectBalancers,
	},

	Route{
		"GetProjectClusters",
		strings.ToUpper("Get"),
		"/api/v1/projects/{project_id}/resources/clusters",
		GetProjectClusters,
	},

	Route{
		"GetProjectDatabases",
		strings.ToUpper("Get"),
		"/api/v1/projects/{project_id}/resources/databases",
		GetProjectDatabases,
	},

	Route{
		"GetProjectDedicatedServers",
		strings.ToUpper("Get"),
		"/api/v1/projects/{project_id}/resources/dedicated",
		GetProjectDedicatedServers,
	},

	Route{
		"GetProjectServers",
		strings.ToUpper("Get"),
		"/api/v1/projects/{project_id}/resources/servers",
		GetProjectServers,
	},

	Route{
		"GetProjectStorages",
		strings.ToUpper("Get"),
		"/api/v1/projects/{project_id}/resources/buckets",
		GetProjectStorages,
	},

	Route{
		"GetProjects",
		strings.ToUpper("Get"),
		"/api/v1/projects",
		GetProjects,
	},

	Route{
		"GetTLD",
		strings.ToUpper("Get"),
		"/api/v1/tlds/{tld_id}",
		GetTLD,
	},

	Route{
		"GetTLDs",
		strings.ToUpper("Get"),
		"/api/v1/tlds",
		GetTLDs,
	},

	Route{
		"TransferResourceToAnotherProject",
		strings.ToUpper("Put"),
		"/api/v1/projects/{project_id}/resources/transfer",
		TransferResourceToAnotherProject,
	},

	Route{
		"UpdateAuthRestrictionsByCountries",
		strings.ToUpper("Post"),
		"/api/v1/auth/access/countries/enabled",
		UpdateAuthRestrictionsByCountries,
	},

	Route{
		"UpdateAuthRestrictionsByIP",
		strings.ToUpper("Post"),
		"/api/v1/auth/access/ips/enabled",
		UpdateAuthRestrictionsByIP,
	},

	Route{
		"UpdateBalancer",
		strings.ToUpper("Patch"),
		"/api/v1/balancers/{balancer_id}",
		UpdateBalancer,
	},

	Route{
		"UpdateBalancerRule",
		strings.ToUpper("Patch"),
		"/api/v1/balancers/{balancer_id}/rules/{rule_id}",
		UpdateBalancerRule,
	},

	Route{
		"UpdateDomainAutoProlongation",
		strings.ToUpper("Patch"),
		"/api/v1/domains/{fqdn}",
		UpdateDomainAutoProlongation,
	},

	Route{
		"UpdateDomainDNSRecord",
		strings.ToUpper("Patch"),
		"/api/v1/domains/{fqdn}/dns-records/{record_id}",
		UpdateDomainDNSRecord,
	},

	Route{
		"UpdateDomainMailInfo",
		strings.ToUpper("Patch"),
		"/api/v1/mail/domains/{domain}/info",
		UpdateDomainMailInfo,
	},

	Route{
		"UpdateDomainNameServers",
		strings.ToUpper("Put"),
		"/api/v1/domains/{fqdn}/name-servers",
		UpdateDomainNameServers,
	},

	Route{
		"UpdateDomainRequest",
		strings.ToUpper("Patch"),
		"/api/v1/domains-requests/{request_id}",
		UpdateDomainRequest,
	},

	Route{
		"UpdateImage",
		strings.ToUpper("Patch"),
		"/api/v1/images/{image_id}",
		UpdateImage,
	},

	Route{
		"UpdateMailQuota",
		strings.ToUpper("Patch"),
		"/api/v1/mail/quota",
		UpdateMailQuota,
	},

	Route{
		"UpdateMailbox",
		strings.ToUpper("Patch"),
		"/api/v1/mail/domains/{domain}/mailboxes/{mailbox}",
		UpdateMailbox,
	},

	Route{
		"UpdateNotificationSettings",
		strings.ToUpper("Patch"),
		"/api/v1/account/notification-settings",
		UpdateNotificationSettings,
	},

	Route{
		"UpdateProject",
		strings.ToUpper("Put"),
		"/api/v1/projects/{project_id}",
		UpdateProject,
	},

	Route{
		"UploadImage",
		strings.ToUpper("Post"),
		"/api/v1/images/{image_id}",
		UploadImage,
	},

	Route{
		"CreateToken",
		strings.ToUpper("Post"),
		"/api/v1/auth/api-keys",
		CreateToken,
	},

	Route{
		"DeleteToken",
		strings.ToUpper("Delete"),
		"/api/v1/auth/api-keys/{token_id}",
		DeleteToken,
	},

	Route{
		"GetTokens",
		strings.ToUpper("Get"),
		"/api/v1/auth/api-keys",
		GetTokens,
	},

	Route{
		"ReissueToken",
		strings.ToUpper("Put"),
		"/api/v1/auth/api-keys/{token_id}",
		ReissueToken,
	},

	Route{
		"UpdateToken",
		strings.ToUpper("Patch"),
		"/api/v1/auth/api-keys/{token_id}",
		UpdateToken,
	},

	Route{
		"AddResourceToGroup",
		strings.ToUpper("Post"),
		"/api/v1/firewall/groups/{group_id}/resources/{resource_id}",
		AddResourceToGroup,
	},

	Route{
		"CreateGroup",
		strings.ToUpper("Post"),
		"/api/v1/firewall/groups",
		CreateGroup,
	},

	Route{
		"CreateGroupRule",
		strings.ToUpper("Post"),
		"/api/v1/firewall/groups/{group_id}/rules",
		CreateGroupRule,
	},

	Route{
		"DeleteGroup",
		strings.ToUpper("Delete"),
		"/api/v1/firewall/groups/{group_id}",
		DeleteGroup,
	},

	Route{
		"DeleteGroupRule",
		strings.ToUpper("Delete"),
		"/api/v1/firewall/groups/{group_id}/rules/{rule_id}",
		DeleteGroupRule,
	},

	Route{
		"DeleteResourceFromGroup",
		strings.ToUpper("Delete"),
		"/api/v1/firewall/groups/{group_id}/resources/{resource_id}",
		DeleteResourceFromGroup,
	},

	Route{
		"GetGroup",
		strings.ToUpper("Get"),
		"/api/v1/firewall/groups/{group_id}",
		GetGroup,
	},

	Route{
		"GetGroupResources",
		strings.ToUpper("Get"),
		"/api/v1/firewall/groups/{group_id}/resources",
		GetGroupResources,
	},

	Route{
		"GetGroupRule",
		strings.ToUpper("Get"),
		"/api/v1/firewall/groups/{group_id}/rules/{rule_id}",
		GetGroupRule,
	},

	Route{
		"GetGroupRules",
		strings.ToUpper("Get"),
		"/api/v1/firewall/groups/{group_id}/rules",
		GetGroupRules,
	},

	Route{
		"GetGroups",
		strings.ToUpper("Get"),
		"/api/v1/firewall/groups",
		GetGroups,
	},

	Route{
		"GetRulesForResource",
		strings.ToUpper("Get"),
		"/api/v1/firewall/service/{resource_type}/{resource_id}",
		GetRulesForResource,
	},

	Route{
		"UpdateGroup",
		strings.ToUpper("Patch"),
		"/api/v1/firewall/groups/{group_id}",
		UpdateGroup,
	},

	Route{
		"UpdateGroupRule",
		strings.ToUpper("Patch"),
		"/api/v1/firewall/groups/{group_id}/rules/{rule_id}",
		UpdateGroupRule,
	},

	Route{
		"CreateCluster",
		strings.ToUpper("Post"),
		"/api/v1/k8s/clusters",
		CreateCluster,
	},

	Route{
		"CreateClusterNodeGroup",
		strings.ToUpper("Post"),
		"/api/v1/k8s/clusters/{cluster_id}/groups",
		CreateClusterNodeGroup,
	},

	Route{
		"DeleteCluster",
		strings.ToUpper("Delete"),
		"/api/v1/k8s/clusters/{cluster_id}",
		DeleteCluster,
	},

	Route{
		"DeleteClusterNode",
		strings.ToUpper("Delete"),
		"/api/v1/k8s/clusters/{cluster_id}/nodes/{node_id}",
		DeleteClusterNode,
	},

	Route{
		"DeleteClusterNodeGroup",
		strings.ToUpper("Delete"),
		"/api/v1/k8s/clusters/{cluster_id}/groups/{group_id}",
		DeleteClusterNodeGroup,
	},

	Route{
		"GetCluster",
		strings.ToUpper("Get"),
		"/api/v1/k8s/clusters/{cluster_id}",
		GetCluster,
	},

	Route{
		"GetClusterKubeconfig",
		strings.ToUpper("Get"),
		"/api/v1/k8s/clusters/{cluster_id}/kubeconfig",
		GetClusterKubeconfig,
	},

	Route{
		"GetClusterNodeGroup",
		strings.ToUpper("Get"),
		"/api/v1/k8s/clusters/{cluster_id}/groups/{group_id}",
		GetClusterNodeGroup,
	},

	Route{
		"GetClusterNodeGroups",
		strings.ToUpper("Get"),
		"/api/v1/k8s/clusters/{cluster_id}/groups",
		GetClusterNodeGroups,
	},

	Route{
		"GetClusterNodes",
		strings.ToUpper("Get"),
		"/api/v1/k8s/clusters/{cluster_id}/nodes",
		GetClusterNodes,
	},

	Route{
		"GetClusterNodesFromGroup",
		strings.ToUpper("Get"),
		"/api/v1/k8s/clusters/{cluster_id}/groups/{group_id}/nodes",
		GetClusterNodesFromGroup,
	},

	Route{
		"GetClusterResources",
		strings.ToUpper("Get"),
		"/api/v1/k8s/clusters/{cluster_id}/resources",
		GetClusterResources,
	},

	Route{
		"GetClusters",
		strings.ToUpper("Get"),
		"/api/v1/k8s/clusters",
		GetClusters,
	},

	Route{
		"GetK8SNetworkDrivers",
		strings.ToUpper("Get"),
		"/api/v1/k8s/network_drivers",
		GetK8SNetworkDrivers,
	},

	Route{
		"GetK8SVersions",
		strings.ToUpper("Get"),
		"/api/v1/k8s/k8s_versions",
		GetK8SVersions,
	},

	Route{
		"GetKubernetesPresets",
		strings.ToUpper("Get"),
		"/api/v1/presets/k8s",
		GetKubernetesPresets,
	},

	Route{
		"IncreaseCountOfNodesInGroup",
		strings.ToUpper("Post"),
		"/api/v1/k8s/clusters/{cluster_id}/groups/{group_id}/nodes",
		IncreaseCountOfNodesInGroup,
	},

	Route{
		"ReduceCountOfNodesInGroup",
		strings.ToUpper("Delete"),
		"/api/v1/k8s/clusters/{cluster_id}/groups/{group_id}/nodes",
		ReduceCountOfNodesInGroup,
	},

	Route{
		"UpdateCluster",
		strings.ToUpper("Patch"),
		"/api/v1/k8s/clusters/{cluster_id}",
		UpdateCluster,
	},

	Route{
		"AddStorageSubdomainCertificate",
		strings.ToUpper("Post"),
		"/api/v1/storages/certificates/generate",
		AddStorageSubdomainCertificate,
	},

	Route{
		"AddStorageSubdomains",
		strings.ToUpper("Post"),
		"/api/v1/storages/buckets/{bucket_id}/subdomains",
		AddStorageSubdomains,
	},

	Route{
		"CopyStorageFile",
		strings.ToUpper("Post"),
		"/api/v1/storages/buckets/{bucket_id}/object-manager/copy",
		CopyStorageFile,
	},

	Route{
		"CreateFolderInStorage",
		strings.ToUpper("Post"),
		"/api/v1/storages/buckets/{bucket_id}/object-manager/mkdir",
		CreateFolderInStorage,
	},

	Route{
		"CreateStorage",
		strings.ToUpper("Post"),
		"/api/v1/storages/buckets",
		CreateStorage,
	},

	Route{
		"DeleteStorage",
		strings.ToUpper("Delete"),
		"/api/v1/storages/buckets/{bucket_id}",
		DeleteStorage,
	},

	Route{
		"DeleteStorageFile",
		strings.ToUpper("Delete"),
		"/api/v1/storages/buckets/{bucket_id}/object-manager/remove",
		DeleteStorageFile,
	},

	Route{
		"DeleteStorageSubdomains",
		strings.ToUpper("Delete"),
		"/api/v1/storages/buckets/{bucket_id}/subdomains",
		DeleteStorageSubdomains,
	},

	Route{
		"GetStorageFilesList",
		strings.ToUpper("Get"),
		"/api/v1/storages/buckets/{bucket_id}/object-manager/list",
		GetStorageFilesList,
	},

	Route{
		"GetStorageSubdomains",
		strings.ToUpper("Get"),
		"/api/v1/storages/buckets/{bucket_id}/subdomains",
		GetStorageSubdomains,
	},

	Route{
		"GetStorageTransferStatus",
		strings.ToUpper("Get"),
		"/api/v1/storages/buckets/{bucket_id}/transfer-status",
		GetStorageTransferStatus,
	},

	Route{
		"GetStorageUsers",
		strings.ToUpper("Get"),
		"/api/v1/storages/users",
		GetStorageUsers,
	},

	Route{
		"GetStorages",
		strings.ToUpper("Get"),
		"/api/v1/storages/buckets",
		GetStorages,
	},

	Route{
		"GetStoragesPresets",
		strings.ToUpper("Get"),
		"/api/v1/presets/storages",
		GetStoragesPresets,
	},

	Route{
		"RenameStorageFile",
		strings.ToUpper("Post"),
		"/api/v1/storages/buckets/{bucket_id}/object-manager/rename",
		RenameStorageFile,
	},

	Route{
		"TransferStorage",
		strings.ToUpper("Post"),
		"/api/v1/storages/transfer",
		TransferStorage,
	},

	Route{
		"UpdateStorage",
		strings.ToUpper("Patch"),
		"/api/v1/storages/buckets/{bucket_id}",
		UpdateStorage,
	},

	Route{
		"UpdateStorageUser",
		strings.ToUpper("Patch"),
		"/api/v1/storages/users/{user_id}",
		UpdateStorageUser,
	},

	Route{
		"UploadFileToStorage",
		strings.ToUpper("Post"),
		"/api/v1/storages/buckets/{bucket_id}/object-manager/upload",
		UploadFileToStorage,
	},

	Route{
		"AddKeyToServer",
		strings.ToUpper("Post"),
		"/api/v1/servers/{server_id}/ssh-keys",
		AddKeyToServer,
	},

	Route{
		"CreateKey",
		strings.ToUpper("Post"),
		"/api/v1/ssh-keys",
		CreateKey,
	},

	Route{
		"DeleteKey",
		strings.ToUpper("Delete"),
		"/api/v1/ssh-keys/{ssh_key_id}",
		DeleteKey,
	},

	Route{
		"DeleteKeyFromServer",
		strings.ToUpper("Delete"),
		"/api/v1/servers/{server_id}/ssh-keys/{ssh_key_id}",
		DeleteKeyFromServer,
	},

	Route{
		"GetKey",
		strings.ToUpper("Get"),
		"/api/v1/ssh-keys/{ssh_key_id}",
		GetKey,
	},

	Route{
		"GetKeys",
		strings.ToUpper("Get"),
		"/api/v1/ssh-keys",
		GetKeys,
	},

	Route{
		"UpdateKey",
		strings.ToUpper("Patch"),
		"/api/v1/ssh-keys/{ssh_key_id}",
		UpdateKey,
	},

	Route{
		"CreateVPC",
		strings.ToUpper("Post"),
		"/api/v2/vpcs",
		CreateVPC,
	},

	Route{
		"DeleteVPC",
		strings.ToUpper("Delete"),
		"/api/v1/vpcs/{vpc_id}",
		DeleteVPC,
	},

	Route{
		"GetVPC",
		strings.ToUpper("Get"),
		"/api/v2/vpcs/{vpc_id}",
		GetVPC,
	},

	Route{
		"GetVPCPorts",
		strings.ToUpper("Get"),
		"/api/v1/vpcs/{vpc_id}/ports",
		GetVPCPorts,
	},

	Route{
		"GetVPCServices",
		strings.ToUpper("Get"),
		"/api/v2/vpcs/{vpc_id}/services",
		GetVPCServices,
	},

	Route{
		"GetVPCs",
		strings.ToUpper("Get"),
		"/api/v2/vpcs",
		GetVPCs,
	},

	Route{
		"UpdateVPCs",
		strings.ToUpper("Patch"),
		"/api/v2/vpcs/{vpc_id}",
		UpdateVPCs,
	},

	Route{
		"AddServerIP",
		strings.ToUpper("Post"),
		"/api/v1/servers/{server_id}/ips",
		AddServerIP,
	},

	Route{
		"CloneServer",
		strings.ToUpper("Post"),
		"/api/v1/servers/{server_id}/clone",
		CloneServer,
	},

	Route{
		"CreateDatabase",
		strings.ToUpper("Post"),
		"/api/v1/dbs",
		CreateDatabase,
	},

	Route{
		"CreateDatabaseBackup",
		strings.ToUpper("Post"),
		"/api/v1/dbs/{db_id}/backups",
		CreateDatabaseBackup,
	},

	Route{
		"CreateDatabaseCluster",
		strings.ToUpper("Post"),
		"/api/v1/databases",
		CreateDatabaseCluster,
	},

	Route{
		"CreateDatabaseInstance",
		strings.ToUpper("Post"),
		"/api/v1/databases/{db_cluster_id}/instances",
		CreateDatabaseInstance,
	},

	Route{
		"CreateDatabaseUser",
		strings.ToUpper("Post"),
		"/api/v1/databases/{db_cluster_id}/admins",
		CreateDatabaseUser,
	},

	Route{
		"CreateDedicatedServer",
		strings.ToUpper("Post"),
		"/api/v1/dedicated-servers",
		CreateDedicatedServer,
	},

	Route{
		"CreateServer",
		strings.ToUpper("Post"),
		"/api/v1/servers",
		CreateServer,
	},

	Route{
		"CreateServerDisk",
		strings.ToUpper("Post"),
		"/api/v1/servers/{server_id}/disks",
		CreateServerDisk,
	},

	Route{
		"CreateServerDiskBackup",
		strings.ToUpper("Post"),
		"/api/v1/servers/{server_id}/disks/{disk_id}/backups",
		CreateServerDiskBackup,
	},

	Route{
		"DeleteDatabase",
		strings.ToUpper("Delete"),
		"/api/v1/dbs/{db_id}",
		DeleteDatabase,
	},

	Route{
		"DeleteDatabaseBackup",
		strings.ToUpper("Delete"),
		"/api/v1/dbs/{db_id}/backups/{backup_id}",
		DeleteDatabaseBackup,
	},

	Route{
		"DeleteDatabaseCluster",
		strings.ToUpper("Delete"),
		"/api/v1/databases/{db_cluster_id}",
		DeleteDatabaseCluster,
	},

	Route{
		"DeleteDatabaseInstance",
		strings.ToUpper("Delete"),
		"/api/v1/databases/{db_cluster_id}/instances/{instance_id}",
		DeleteDatabaseInstance,
	},

	Route{
		"DeleteDatabaseUser",
		strings.ToUpper("Delete"),
		"/api/v1/databases/{db_cluster_id}/admins/{admin_id}",
		DeleteDatabaseUser,
	},

	Route{
		"DeleteDedicatedServer",
		strings.ToUpper("Delete"),
		"/api/v1/dedicated-servers/{dedicated_id}",
		DeleteDedicatedServer,
	},

	Route{
		"DeleteServer",
		strings.ToUpper("Delete"),
		"/api/v1/servers/{server_id}",
		DeleteServer,
	},

	Route{
		"DeleteServerDisk",
		strings.ToUpper("Delete"),
		"/api/v1/servers/{server_id}/disks/{disk_id}",
		DeleteServerDisk,
	},

	Route{
		"DeleteServerDiskBackup",
		strings.ToUpper("Delete"),
		"/api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id}",
		DeleteServerDiskBackup,
	},

	Route{
		"DeleteServerIP",
		strings.ToUpper("Delete"),
		"/api/v1/servers/{server_id}/ips",
		DeleteServerIP,
	},

	Route{
		"GetConfigurators",
		strings.ToUpper("Get"),
		"/api/v1/configurator/servers",
		GetConfigurators,
	},

	Route{
		"GetDatabase",
		strings.ToUpper("Get"),
		"/api/v1/dbs/{db_id}",
		GetDatabase,
	},

	Route{
		"GetDatabaseAutoBackupsSettings",
		strings.ToUpper("Get"),
		"/api/v1/dbs/{db_id}/auto-backups",
		GetDatabaseAutoBackupsSettings,
	},

	Route{
		"GetDatabaseBackup",
		strings.ToUpper("Get"),
		"/api/v1/dbs/{db_id}/backups/{backup_id}",
		GetDatabaseBackup,
	},

	Route{
		"GetDatabaseBackups",
		strings.ToUpper("Get"),
		"/api/v1/dbs/{db_id}/backups",
		GetDatabaseBackups,
	},

	Route{
		"GetDatabaseCluster",
		strings.ToUpper("Get"),
		"/api/v1/databases/{db_cluster_id}",
		GetDatabaseCluster,
	},

	Route{
		"GetDatabaseClusters",
		strings.ToUpper("Get"),
		"/api/v1/databases",
		GetDatabaseClusters,
	},

	Route{
		"GetDatabaseInstance",
		strings.ToUpper("Get"),
		"/api/v1/databases/{db_cluster_id}/instances/{instance_id}",
		GetDatabaseInstance,
	},

	Route{
		"GetDatabaseInstances",
		strings.ToUpper("Get"),
		"/api/v1/databases/{db_cluster_id}/instances",
		GetDatabaseInstances,
	},

	Route{
		"GetDatabaseUser",
		strings.ToUpper("Get"),
		"/api/v1/databases/{db_cluster_id}/admins/{admin_id}",
		GetDatabaseUser,
	},

	Route{
		"GetDatabaseUsers",
		strings.ToUpper("Get"),
		"/api/v1/databases/{db_cluster_id}/admins",
		GetDatabaseUsers,
	},

	Route{
		"GetDatabases",
		strings.ToUpper("Get"),
		"/api/v1/dbs",
		GetDatabases,
	},

	Route{
		"GetDatabasesPresets",
		strings.ToUpper("Get"),
		"/api/v1/presets/dbs",
		GetDatabasesPresets,
	},

	Route{
		"GetDedicatedServer",
		strings.ToUpper("Get"),
		"/api/v1/dedicated-servers/{dedicated_id}",
		GetDedicatedServer,
	},

	Route{
		"GetDedicatedServerPresetAdditionalServices",
		strings.ToUpper("Get"),
		"/api/v1/presets/dedicated-servers/{preset_id}/additional-services",
		GetDedicatedServerPresetAdditionalServices,
	},

	Route{
		"GetDedicatedServers",
		strings.ToUpper("Get"),
		"/api/v1/dedicated-servers",
		GetDedicatedServers,
	},

	Route{
		"GetDedicatedServersPresets",
		strings.ToUpper("Get"),
		"/api/v1/presets/dedicated-servers",
		GetDedicatedServersPresets,
	},

	Route{
		"GetOsList",
		strings.ToUpper("Get"),
		"/api/v1/os/servers",
		GetOsList,
	},

	Route{
		"GetServer",
		strings.ToUpper("Get"),
		"/api/v1/servers/{server_id}",
		GetServer,
	},

	Route{
		"GetServerDisk",
		strings.ToUpper("Get"),
		"/api/v1/servers/{server_id}/disks/{disk_id}",
		GetServerDisk,
	},

	Route{
		"GetServerDiskAutoBackupSettings",
		strings.ToUpper("Get"),
		"/api/v1/servers/{server_id}/disks/{disk_id}/auto-backups",
		GetServerDiskAutoBackupSettings,
	},

	Route{
		"GetServerDiskBackup",
		strings.ToUpper("Get"),
		"/api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id}",
		GetServerDiskBackup,
	},

	Route{
		"GetServerDiskBackups",
		strings.ToUpper("Get"),
		"/api/v1/servers/{server_id}/disks/{disk_id}/backups",
		GetServerDiskBackups,
	},

	Route{
		"GetServerDisks",
		strings.ToUpper("Get"),
		"/api/v1/servers/{server_id}/disks",
		GetServerDisks,
	},

	Route{
		"GetServerIPs",
		strings.ToUpper("Get"),
		"/api/v1/servers/{server_id}/ips",
		GetServerIPs,
	},

	Route{
		"GetServerLogs",
		strings.ToUpper("Get"),
		"/api/v1/servers/{server_id}/logs",
		GetServerLogs,
	},

	Route{
		"GetServerStatistics",
		strings.ToUpper("Get"),
		"/api/v1/servers/{server_id}/statistics",
		GetServerStatistics,
	},

	Route{
		"GetServers",
		strings.ToUpper("Get"),
		"/api/v1/servers",
		GetServers,
	},

	Route{
		"GetServersPresets",
		strings.ToUpper("Get"),
		"/api/v1/presets/servers",
		GetServersPresets,
	},

	Route{
		"GetSoftware",
		strings.ToUpper("Get"),
		"/api/v1/software/servers",
		GetSoftware,
	},

	Route{
		"ImageUnmountAndServerReload",
		strings.ToUpper("Post"),
		"/api/v1/servers/{server_id}/image-unmount",
		ImageUnmountAndServerReload,
	},

	Route{
		"PerformActionOnBackup",
		strings.ToUpper("Post"),
		"/api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id}/action",
		PerformActionOnBackup,
	},

	Route{
		"PerformActionOnServer",
		strings.ToUpper("Post"),
		"/api/v1/servers/{server_id}/action",
		PerformActionOnServer,
	},

	Route{
		"RestoreDatabaseFromBackup",
		strings.ToUpper("Put"),
		"/api/v1/dbs/{db_id}/backups/{backup_id}",
		RestoreDatabaseFromBackup,
	},

	Route{
		"UpdateDatabase",
		strings.ToUpper("Patch"),
		"/api/v1/dbs/{db_id}",
		UpdateDatabase,
	},

	Route{
		"UpdateDatabaseAutoBackupsSettings",
		strings.ToUpper("Patch"),
		"/api/v1/dbs/{db_id}/auto-backups",
		UpdateDatabaseAutoBackupsSettings,
	},

	Route{
		"UpdateDatabaseCluster",
		strings.ToUpper("Patch"),
		"/api/v1/databases/{db_cluster_id}",
		UpdateDatabaseCluster,
	},

	Route{
		"UpdateDatabaseInstance",
		strings.ToUpper("Patch"),
		"/api/v1/databases/{db_cluster_id}/instances/{instance_id}",
		UpdateDatabaseInstance,
	},

	Route{
		"UpdateDatabaseUser",
		strings.ToUpper("Patch"),
		"/api/v1/databases/{db_cluster_id}/admins/{admin_id}",
		UpdateDatabaseUser,
	},

	Route{
		"UpdateDedicatedServer",
		strings.ToUpper("Patch"),
		"/api/v1/dedicated-servers/{dedicated_id}",
		UpdateDedicatedServer,
	},

	Route{
		"UpdateServer",
		strings.ToUpper("Patch"),
		"/api/v1/servers/{server_id}",
		UpdateServer,
	},

	Route{
		"UpdateServerDisk",
		strings.ToUpper("Patch"),
		"/api/v1/servers/{server_id}/disks/{disk_id}",
		UpdateServerDisk,
	},

	Route{
		"UpdateServerDiskAutoBackupSettings",
		strings.ToUpper("Patch"),
		"/api/v1/servers/{server_id}/disks/{disk_id}/auto-backups",
		UpdateServerDiskAutoBackupSettings,
	},

	Route{
		"UpdateServerDiskBackup",
		strings.ToUpper("Patch"),
		"/api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id}",
		UpdateServerDiskBackup,
	},

	Route{
		"UpdateServerIP",
		strings.ToUpper("Patch"),
		"/api/v1/servers/{server_id}/ips",
		UpdateServerIP,
	},

	Route{
		"UpdateServerNAT",
		strings.ToUpper("Patch"),
		"/api/v1/servers/{server_id}/local-networks/nat-mode",
		UpdateServerNAT,
	},

	Route{
		"UpdateServerOSBootMode",
		strings.ToUpper("Post"),
		"/api/v1/servers/{server_id}/boot-mode",
		UpdateServerOSBootMode,
	},
}
