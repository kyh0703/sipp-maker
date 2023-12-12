ACK sip:[field1]@[remote_ip]:[remote_port] SIP/2.0
Via: SIP/2.0/[transport] [local_ip]:[local_port];branch=[branch]
Max-Forwards: 70
From: <tel:[field0]@[local_ip]:[local_port]>;tag=[call_number]
To: <tel:[field0]@[remote_ip]:[remote_port]>[peer_tag_param]
Call-ID: [call_id]
CSeq: 1 ACK
Contact: sip:[field0]@[local_ip]:[local_port]
Subject: Sipp Scenario Test
Content-Length: 0