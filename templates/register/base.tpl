<?xml version="1.0" encoding="ISO-8859-1" ?>
<!DOCTYPE scenario SYSTEM "sipp.dtd">

<scenario name="sipp-maker Register Scenario">

<send crlf="true" retrans="500">
<![CDATA[
{{ if eq .UriFormat "sip" }}
{{- template "register_sip.tpl" -}}
{{ else }}
{{- template "register_tel.tpl" -}}
{{ end }}
]]>
</send>

<recv response="100" optional="true" />
<recv response="401" auth="true" rtd="true" />

<![CDATA[
{{ if eq .UriFormat "sip" }}
{{- template "auth_sip.tpl" -}}
{{ else }}
{{- template "auth_tel.tpl" -}}
{{ end }}
]]>

<ResponseTimeRepartition value="10, 20"/>
<CallLengthRepartition value="10"/>
</scenario>
