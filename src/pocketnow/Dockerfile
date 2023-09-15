# use this to debug docker build . --no-cache --progress=plain 

FROM --platform=$BUILDPLATFORM  mcr.microsoft.com/dotnet/sdk:8.0-jammy AS build
ARG TARGETARCH

WORKDIR /source

# copy csproj and restore as distinct layers
COPY pocketnow.sln .
# dest folders aren't created automatically
RUN mkdir pocketnow.lib
RUN mkdir pocketnow.service
RUN mkdir pocketnow.Tests
COPY pocketnow.lib/pocketnow.lib.csproj pocketnow.lib/.
COPY pocketnow.service/*.csproj pocketnow.service/. 
COPY pocketnow.Tests/*.csproj pocketnow.Tests/.
# verify things are where expected
#RUN ls -al
#RUN ls -al pocketnow.lib
RUN dotnet restore -a $TARGETARCH pocketnow.sln


## copy everything else and build app
#COPY . ./
#WORKDIR /source/
###RUN dotnet publish -c release -o /app
##RUN dotnet publish -c release -o /app --no-restore
#RUN dotnet restore pocketnow.sln
##RUN dotnet publish -c release -r linux-x64 pocketnow.sln --self-contained
#RUN dotnet publish -c release -o /app pocketnow.sln
#
## final stage/image
##FROM mcr.microsoft.com/dotnet/aspnet:6.0
#FROM mcr.microsoft.com/dotnet/aspnet:8.0.0-preview.7-bookworm-slim
##FROM mcr.microsoft.com/dotnet/aspnet:7.0
#WORKDIR /app
#ENV ASPNETCORE_URLS=http://0.0.0.0:5000;
#EXPOSE 5000
#COPY --from=build /app ./
#ENTRYPOINT ["dotnet", "pocketnow.dll"]
