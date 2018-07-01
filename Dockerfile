FROM telegrammessenger/proxy
COPY mtproto_stats .
RUN echo "$(tail -n +2 run.sh)" > run.sh && echo '#!/bin/bash\n./mtproto_stats & disown' | cat - run.sh > temp && mv temp run.sh
CMD [ "/bin/sh", "-c", "/bin/bash /run.sh"] 