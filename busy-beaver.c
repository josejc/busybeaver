/* DFB's busy beaver (initially  the 3 card case */
// Info -> https://en.wikipedia.org/wiki/Busy_beaver

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <time.h>

#define UPBTAPE 25	// should be plenty since SIGMA(6) is way out of reach ? 
#define LWBTAPE 1
#define FALSE 0
#define TRUE 1

//typedef struct {int zeroinpt; int zeroactions[3]; int oneinpt; int oneactions[3];} CARD;

//trivial single-card examples
// loops off to the right
char * tm1a [2] [2] = {"0000", "1000", "0111", "1111"};
// loops off to the left
char * tm1b [2] [2] = {"0000", "1000", "0101", "1101"};
//prints a single 1 and moves to the left
char * tm1c [2] [2] = {"0000", "1000", "0100", "1000"};
//prints a single 1 and moves to the right
char * tm1d [2] [2] = {"0000", "1000", "0110", "1000"};
// simple example of 2-state with optimal score of 4.
char * tm2 [3] [2] = {"0000", "1000", "0112", "1102", "1101", "1110"};
//classic 3-state BB producing a score of 6
char * tm3a [4] [2] = {"0000", "1000", "0102", "1113", "0111", "1102", "0112", "1100"};
//another 3-state BB for sigma ?
char * tm3aa [4] [2] = {"0000", "1000", "0112", "1110", "0013", "1112", "0103", "1101"};
//permutation of tm3a which never terminates
char * tm3b [4] [2] = {"0000", "1000", "0112", "1100", "0111", "1102", "0102", "1113"};
//another permutation of tm3a which terminates but has a score of only 2
char * tm3c [4] [2] = {"0000", "1000", "0102", "1113", "0112", "1100", "0111", "1102"};
//classic 4-state tm producing a score of 23 in 107 moves
char * tm4 [5] [2] = {"0000", "1000", "0112", "1102", "0101", "1003", "0110", "1104", "0114", "1011" };

// first bound gives us space for up to 4-state BB machines. row zero is
// unused so that card numbering can start at 1. Card zero effectively represents the halted state


char tape [25] = "0000000000000000000000000";
//int tapelen = strlen(tape);
int tapelen = UPBTAPE - LWBTAPE - 1;
int  headpos, currtcard = 1, steps = 0;

void printtape ( )
{
// print out the tape and on the line below it print out a ^ to show where head pos is for the
// next read/write
    int i;
    printf("%d: ", steps);
    for (i = 0; i < tapelen; i++)
        printf ("%c", tape[i] );
    putchar ('\n');

    printf("%d: ", steps);
    for( i=0 ; i < headpos ; i++)
        printf(" ") ;
    putchar('^');
    putchar('\n');
}


int runtm (char* thistm [] [2], int n)
{
//run the tm and make stepwise changes to tape.
// "n" is the number of cards (excluding the halt card) in this BB

//local variables
    char currtchar, chtowrite;
    int i, hdstate, nextcard, incdec;


//print out initial deck of cards
    for (i = 0; i <= n; i++)
    {
//	if (i>0 && strcmp(thistm[i][0],"0000") == 0 ) break;
        printf ("Card %d is %s  %s\n", i, thistm[i][0], thistm[i][1]);
    }
    headpos = tapelen / 2;
    printf("\nInitial tape is shown below: Note that head position at each step is marked by ^\n\n");
    printtape();

    printf("\nGame started\n");

    do {
        currtchar = tape[headpos];
        steps++;
        hdstate = currtchar - '0'; //denotes whether a 0 or a 1 has just been read
        chtowrite = thistm[currtcard][hdstate][1];
        char shift = thistm[currtcard][hdstate][2];
        nextcard =  thistm[currtcard][hdstate][3] - '0';
        int incdec = (shift == '0') ? -1 : +1; // shift headpos left or right after next write

//printf ("starting char from C1 column 0 is %c \n", currtchar);
//sanity check
        if ((headpos < LWBTAPE-1) || (headpos == UPBTAPE - 2))
        {
            printf ("tape too short and/or TM may be looping\n");
            return(1);
        }
#ifdef DEBUG
        printf ("char to print is %c \n", chtowrite);
        printf ("shiftchar is %c \n", shift);
        printf ("nextcard is  %d \n", nextcard);
        printf ("headpos is  now %d \n", headpos);
#endif
        tape[headpos] = chtowrite;
        printtape();
        headpos = headpos + incdec;
        currtcard = nextcard;
    } while (currtcard != 0);
}

int main (char argc, char *argv[])
{
    int i, j , retcode, ncards, score =0;
    char **currentgame;
// Uncomment the statements below for the game you want to be played.
// Then recompile and execute. Ideally all of this needs to be re-programmed to prompt the user
// at run-time to choose the game to be played.

//currentgame = (char **) tm1a; ncards = 1;
//currentgame = (char **) tm1b; ncards = 1;
//currentgame = (char **) tm1c; ncards = 1;
//currentgame = (char **) tm1d; ncards = 1;
//currentgame = (char **) tm2; ncards = 2;
//currentgame = (char **) tm3a; ncards = 3;
//currentgame = (char **) tm3aa; ncards = 3;
//currentgame = (char **) tm3b; ncards = 3;
//currentgame = (char **) tm3c; ncards = 3;
    currentgame = (char **) tm4;
    ncards = 4;

// now run the appropriate BB Turing Machine

    retcode = runtm ((char *(*)[2]) currentgame, ncards);
    if (retcode == 0) printf("The TM has halted\n");

    printf ("\nFinal tape config is:\n\n");
    printtape();

    if (retcode != 1)
    {
        for (i == LWBTAPE; i<= UPBTAPE; i++)
            if (tape[i] == '1') score++;
        printf("\nsteps = %d    score = %d\n", steps, score);
    }
}
