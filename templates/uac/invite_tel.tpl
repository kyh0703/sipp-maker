INVITE sip:[field1]@[remote_ip]:[remote_port] SIP/2.0
Via: SIP/2.0/[transport] [local_ip]:[local_port];branch=[branch]
Max-Forwards: 70
From: <tel:[field0]@[local_ip]:[local_port]>;tag=[call_number]
To: <tel:[field1]@[remote_ip]>
Call-ID: [call_id]
CSeq: 1 INVITE
Contact: sip:[field0]@[local_ip]:[local_port]
Subject: Sipp Scenario Test
Content-Type: application/sdp
Content-Length: [len]