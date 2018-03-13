Lepsius
=======

[Karl Richard Lepsius](https://en.wikipedia.org/wiki/Karl_Richard_Lepsius)
was the first modern reader of the
[book of dead](https://en.wikipedia.org/wiki/Book_of_the_Dead).

Lepsius convert logs into events.

Example
-------

    ---
    input:
      - journald:
          matches:
            SYSLOG_IDENTIFIER: sshd
    output:
      - stdout:

License
-------

3 terms BSD License, ©2016 Mathieu Lecarme
