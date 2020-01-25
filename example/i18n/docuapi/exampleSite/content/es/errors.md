---
weight: 20
title: Errores

---

# Errores

<aside class = "notice"> Esta sección error se almacena en un archivo separado, errors.md. DocuAPI le permite dividir la documentación de una sola página en tantos archivos como sea necesario. Los archivos se incluyen en el orden de las páginas Hugo defecto. Una forma de ordenar las páginas es mediante el establecimiento de la página `weight` en el asunto principal. Las páginas con menor peso se enumeran en primer lugar. </aside>

La API Kittn utiliza los siguientes códigos de error:


Código de error | Sentido
---------- | -------
400 | Bad Request - Solicitud de información chupa
401 | No autorizada - Tu clave de API está mal
403 | Prohibida - El gatito solicitado está oculto sólo para administradores
404 | No se ha encontrado - El gatito especificado no se pudo encontrar
405 | Método no permitido - que intentó acceder a un gatito con un método válido
406 | No Aceptable - Ha solicitado un formato que no es JSON
410 | Ido - El gatito solicitado ha sido eliminado de nuestros servidores
418 | Soy una tetera
429 | Demasiadas peticiones - que está solicitando demasiados gatitos! ¡Ve más despacio!
500 | Error interno del servidor - Tuvimos un problema con nuestro servidor. Inténtelo de nuevo más tarde.
503 | Servicio no disponible - Estamos fuera de línea para temporarially hecho mantenimiento. Por favor, inténtelo de nuevo más tarde.
