package template

var DOCKER_TEMPLATES = []Item{
	{Id: 0, Title: "mssql", Desc: "Install the SQL Server container image", Command: "docker", CommandArgs: " run -e 'ACCEPT_EULA=1' -e 'MSSQL_SA_PASSWORD=p@ssw0rd' -e 'MSSQL_PID=Developer' -e 'MSSQL_USER=SA' -p 1433:1433 -d --name=sql mcr.microsoft.com/azure-sql-edge"},
}
