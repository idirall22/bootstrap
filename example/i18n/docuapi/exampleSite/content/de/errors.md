---
weight: 20
title: Fehler

---

# Errors

<aside class = "notice"> Dieser Fehler Abschnitt in einer separaten Datei gespeichert ist, errors.md. DocuAPI können Sie die einzelnen Seite Dokumentation in so viele Dateien aufgeteilt je nach Bedarf. Die Dateien werden in der Standard-Hugo-Seite, um enthalten. Eine Möglichkeit, die Seiten zu bestellen, indem Sie die Seite `weight` in der Titelei Einstellung. Seiten mit geringerem Gewicht werden zuerst aufgeführt werden. </aside>

Die Kittn API verwendet die folgenden Fehlercodes:


Fehlercode | Bedeutung
---------- | -------
400 | Bad Request - Ihre Anfrage saugt
401 | Nicht autorisierte - Ihr API-Schlüssel ist falsch
403 | nur Das Kätzchen Verlangte ist für Administratoren versteckt - Verboten
404 | Nicht gefunden - Der angegebene Kätzchen konnte nicht gefunden werden
405 | Methode nicht erlaubt - Sie für den Zugriff versucht, ein Kätzchen mit einem ungültigen Verfahren
406 | Nicht akzeptabel - Sie baten um ein Format, das nicht json ist
410 | Gegangen - angefordert Das Kätzchen wurde von unseren Servern entfernt
418 | Ich bin eine Teekanne
429 | Zu viele Anfragen - Sie anfordern zu viele Kätzchen! Langsamer!
500 | Internal Server Error - Wir hatten ein Problem mit unserem Server. Versuchen Sie es später noch einmal.
503 | Nicht verfügbar Service - Wir sind temporarially offline für maintanance. Bitte versuchen Sie es später noch einmal.
