# https://hub.docker.com/_/microsoft-dotnet
FROM mcr.microsoft.com/dotnet/sdk:8.0.100-preview.7-bookworm-slim AS build
WORKDIR /source

# copy csproj and restore as distinct layers
# COPY *.csproj ./
# RUN dotnet restore

# copy everything else and build app
COPY . ./
WORKDIR /source/
##RUN dotnet publish -c release -o /app
#RUN dotnet publish -c release -o /app --no-restore
RUN dotnet restore pocketnow.sln
RUN dotnet publish -c release -r linux-x64 pocketnow.sln --self-contained

# final stage/image
#FROM mcr.microsoft.com/dotnet/aspnet:6.0
FROM mcr.microsoft.com/dotnet/aspnet:8.0.0-preview.7-bookworm-slim
#FROM mcr.microsoft.com/dotnet/aspnet:7.0
WORKDIR /app
ENV ASPNETCORE_URLS=http://0.0.0.0:5000;
EXPOSE 5000
COPY --from=build /app ./
ENTRYPOINT ["dotnet", "pocketnow.dll"]
