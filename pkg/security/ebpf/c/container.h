#ifndef _CONTAINER_H_
#define _CONTAINER_H_

static __attribute__((always_inline)) u32 copy_container_id(const char src[CONTAINER_ID_LEN], char dst[CONTAINER_ID_LEN]) {
    if (src[0] == 0) {
        return 0;
    }

#pragma unroll
    for (int i = 0; i < CONTAINER_ID_LEN; i++)
    {
        if (src[i] == 0) {
            break;
        }

        dst[i] = src[i];
    }
    return CONTAINER_ID_LEN;
}

#endif
