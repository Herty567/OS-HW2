# OS-HW2
The design of the code is a Ticket Lock that ensures fairness through a turn-based system and theres the
Compare-And-Swap Spin Lock which keeps attempting to aquire the locks.
I ran the code through Visual Studio using the Go Plugin. Just have to put the code in there and hit run.
This will cause different benchmanrks to be shown across different threads.
My findings were that the Ticket Lock performed much more consistently under higher work loads since it 
has fair ordering while the CAS Lock had increased waiting times cause of it's busy-waiting behavior.
The libraries used were fmt, sync, sync/atomic, and time.
