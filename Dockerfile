FROM telegrammessenger/proxy
COPY mtproto_proxy_stat .
RUN echo "$(tail -n +2 run.sh)" > run.sh && echo '#!/bin/bash\n./mtproto_proxy_stat & disown' | cat - run.sh > temp && mv temp run.sh
CMD [ "/bin/sh", "-c", "/bin/bash /run.sh"] 