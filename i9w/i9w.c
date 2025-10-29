#include <math.h>
#include <stdio.h>

#include "vendor/libtai/caltime.h"
#include "vendor/libtai/leapsecs.h"
#include "vendor/libtai/taia.h"

int main(int argc,char *argv[]) {
  struct caltime ref_ct;
  struct taia ref;
  struct taia t;

  if (leapsecs_init() == -1) {
    fprintf(stderr, "fatal: unable to init leapsecs\n");
    return 111;
  }

  if (!caltime_scan("1970-01-01 00:00:00 +0000", &ref_ct)) {
    fprintf(stderr, "fatal: failed to scan reference time\n");
    return 111;
  }

  // We need to zero the nanosecond and attosecond field of the reference ourselves.
  taia_sub(&ref, &ref, &ref);
  caltime_tai(&ref_ct, &ref.sec);

  taia_now(&t);
  leapsecs_add(&t.sec, 1);
  // Both tai_now and caltime_scan add a constant offset to seconds.
  // Subtracting the reference removes the offset, leaving elapsed seconds.
  taia_sub(&t, &t, &ref);

  // leapsecs_add does not include the 10 leap seconds pre-added to TAI.
  double millisec = (taia_approx(&t) + 10) * 1000;
  double beats = fmod(millisec / 86400, 1000);
  printf(":%06.2f\n", beats);

  return 0;
}
