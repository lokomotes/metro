FROM golang:1.11 AS golang

COPY . /go/.

RUN go build /go/src/github.com/lokomotes/metro/cmd/metro-server/main.go



FROM mcr.microsoft.com/windows/servercore:1803

# https://github.com/docker/for-win/issues/1976#issuecomment-463044772
# TODO: not working on nanoserver
SHELL ["powershell", "-Command", "$ErrorActionPreference = 'Stop'; $ProgressPreference = 'SilentlyContinue';"]

RUN Enable-LocalUser Administrator; \
    SecEdit.exe /export /cfg secpol.cfg; \
    (Get-Content secpol.cfg).Replace('PasswordComplexity = 1', 'PasswordComplexity = 0') | \
    Out-File secpol.cfg; \
    SecEdit.exe /configure /db C:\Windows\Security\Local.sdb /cfg secpol.cfg /areas SECURITYPOLICY; \
    Remove-Item secpol.cfg; \
    Set-LocalUser Administrator -Password (New-Object SecureString)

COPY --from=golang /gopath/main.exe /metro-server.exe

ENTRYPOINT [ "/metro-server" ]
