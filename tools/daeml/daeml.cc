// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// copy from: https://github.com/meili/TeamTalk/blob/master/server/src/tools/daeml.cpp
//

#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <signal.h>
#include <errno.h>

/* closeall() -- close all FDs >= a specified value */

void closeall(int fd) {
    int fdlimit = sysconf(_SC_OPEN_MAX);

    if (fdlimit > 128) {
        fdlimit = 128;
    }

    while (fd < fdlimit)
        close(fd++);
}


/* daemon() - detach process from user and disappear into the background
 * returns -1 on failure, but you can't do much except exit in that case
 * since we may already have forked. This is based on the BSD version,
 * so the caller is responsible for things like the umask, etc.
 */

int daemon(int nochdir, int noclose, int asroot) {
    switch (fork()) {
        case 0:  break;
        case -1: return -1;
        default: _exit(0);          /* exit the original process */
    }

    if (setsid() < 0)               /* shoudn't fail */
        return -1;

    if ( !asroot && (setuid(1) < 0) )              /* shoudn't fail */
        return -1;

    /* dyke out this switch if you want to acquire a control tty in */
    /* the future -- not normally advisable for daemons */

    switch (fork()) {
        case 0:  break;
        case -1: return -1;
        default: _exit(0);
    }

    if (!nochdir)
        chdir("/");

    if (!noclose) {
        closeall(0);

        int fd = open("/dev/null", O_RDWR, 0);
        if (fd < 0) {
            printf("open failed, errno=%d\n", errno);
            return -1;
        }

        dup2(fd, STDIN_FILENO);
        dup2(fd, STDOUT_FILENO);
        dup2(fd, STDERR_FILENO);
    }

    return 0;
}

#define TEXT(a) a
void PrintUsage(char* name) {
    printf(
            TEXT("\n ----- \n\n")
            TEXT("Usage:\n")
            TEXT("   	%s program_name \n\n")
            TEXT("Where:\n")
            TEXT("   	%s - Name of this Daemon loader.\n")
            TEXT("   	program_name - Name (including path) of the program you want to load as daemon.\n\n")
            TEXT("Example:\n")
            TEXT("   	%s ./atprcmgr - Launch program 'atprcmgr' in current directory as daemon. \n\n\n\n"),
            name, name, name
            );
}

int main(int argc, char* argv[]) {
    printf(
           TEXT("\n")
           TEXT("Daemon loader\n")
           TEXT("- Launch specified program as daemon.\n")
           //TEXT("- Require root privilege to launch successfully.\n\n\n")
           );

    if (argc < 2) {
        printf("* Missing parameter : daemon program name not specified!\n");
        PrintUsage(argv[0]);
        exit(0);
    }

    printf("- Loading %s as daemon, please wait ......\n\n\n", argv[1]);

    if (daemon(1, 0, 1) >= 0) {
        signal(SIGCHLD, SIG_IGN);

        //execl(argv[1], argv[1], NULL);
        execv(argv[1], argv + 1);
        printf("! Excute daemon programm %s failed. \n", argv[1]);

        exit(0);
    }

    printf("! Create daemon error. Please check if you have 'root' privilege. \n");
    return 0;
}
