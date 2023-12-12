<?xml version="1.0" encoding="ISO-8859-1" ?>
<!DOCTYPE scenario SYSTEM "sipp.dtd">

<scenario name="sipp-maker UAS Scenario">

<recv request="REGISTER" optional="true" crlf="true" next="LABEL_RES">
<action>
<ereg regexp=": *(.*)" search_in="hdr" header="From" start_line="true" assign_to="_,them"/>
<ereg regexp=": *(.*)" search_in="hdr" header="To" start_line="true" assign_to="_,us"/>
</action>
</recv>

<recv request="OPTIONS" optional="true" crlf="true" next="LABEL_RES">
<action>
<ereg regexp=": *(.*)" search_in="hdr" header="From" start_line="true" assign_to="_,them"/>
<ereg regexp=": *(.*)" search_in="hdr" header="To" start_line="true" assign_to="_,us"/>
</action>
</recv>

</scenario>
