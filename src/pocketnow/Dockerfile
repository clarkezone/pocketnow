# https://hub.docker.com/_/microsoft-dotnet
# FROM mcr.microsoft.com/dotnet/sdk:8.0.100-preview.6-bookworm-slim AS build
from mcr.microsoft.com/dotnet/sdk:8.0.100-preview.6-bookworm-slim@sha256:bf7142071c0a2f4c355fa95522c53086f7e4c5864cf90c4b127cedcad2166dd4 AS build
WORKDIR /source

# copy csproj and restore as distinct layers
COPY *.csproj ./
RUN dotnet restore

# copy everything else and build app
COPY . ./
WORKDIR /source/
RUN dotnet publish -c release -o /app --no-restore

# final stage/image
#FROM mcr.microsoft.com/dotnet/aspnet:6.0
#FROM mcr.microsoft.com/dotnet/aspnet:8.0.0-bullseye-slim-arm64v8
FROM mcr.microsoft.com/dotnet/aspnet:8.0-preview-alpine
#FROM mcr.microsoft.com/dotnet/aspnet:7.0
WORKDIR /app
ENV ASPNETCORE_URLS=http://0.0.0.0:5000;
EXPOSE 5000
COPY --from=build /app ./
ENTRYPOINT ["dotnet", "pocketnow.dll"]
