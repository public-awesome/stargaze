FROM faddat/sos-lite

# TODO: motd with logo & banner text
# COPY motd /etc/motd

RUN sed -i -e "s/^CheckSpace/#!!!CheckSpace/g" /etc/pacman.conf && \
                pacman -Syyu --noconfirm rust && \
                echo 'docker run -d -v starsd:starsd faddat/starsd /usr/bin/starsd init' >> /usr/local/bin/firstboot.sh && \
                echo 'docker run -d -v starsd:/starsd --restart=unless-stopped --name starsd faddat/starsd /usr/bin/starsd --home /starsd' >> /usr/local/bin/firstboot.sh && \
                sed -i -e "s/^#!!!CheckSpace/CheckSpace/g" /etc/pacman.conf


