<?xml version="1.0" encoding="ISO-8859-1" ?>
<!DOCTYPE scenario SYSTEM "sipp.dtd">

<scenario name="sipp-maker UAC Scenario">

<send retrans="500">
<![CDATA[
{{ if eq .UriFormat "sip" }}
{{- template "invite_sip.tpl" -}}
{{ else }}
{{- template "invite_tel.tpl" -}}
{{ end }}

{{ if .IsSdp }}
{{- template "sdp.tpl" -}}
{{ end }}
]]>
</send>

<recv response="100" optional="true">
<recv response="180" optional="true">
<recv response="183" optional="true">
<recv response="200" rtd="true">

<send>
<![CDATA[
{{ if eq .UriFormat "sip" }}
{{- template "ack_sip.tpl" -}}
{{ else }}
{{- template "ack_tel.tpl" -}}
{{ end}}
]]>
</send>

<ResponseTimeRepartition value="10, 20, 30, 40, 50, 100, 150, 200"/>
<CallLengthRepartition value="10, 50, 100, 500, 1000, 5000, 10000"/>
</scenario>
