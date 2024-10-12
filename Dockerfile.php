# Imagen base para PHP
FROM php:7.4-apache

# Copiar el contenido de la carpeta php en la carpeta /var/www/html
COPY php/index.php /var/www/html/

# Exponer el puerto 8080 
EXPOSE 8080

# Configurar el archivo de configuración de Apache para que use el puerto 8080 en lugar de 80
RUN sed -i 's/80/8080/g' /etc/apache2/sites-available/000-default.conf /etc/apache2/ports.conf

# Ejecutar el comando de inicio de Apache
CMD ["apache2-foreground"]