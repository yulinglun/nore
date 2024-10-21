package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tharsis/evmos/app"
	cmdcfg "github.com/tharsis/evmos/cmd/config"
)

func main() {
	setupConfig()
	cmdcfg.RegisterDenoms()

	rootCmd, _ := NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}

func setupConfig() {
	// set the address prefixes
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	cmdcfg.SetBip44CoinType(config)
	config.Seal()
}
void jforward_list_remove(JForwardList *jlist, int value) {
  struct SingleNode *current = jlist->head;
  struct SingleNode *prev = NULL;

  while (current) {
    if (current->data == value) {
      if (prev) {
        prev->next = current->next;
        jlist->tail = current == jlist->tail ? prev : jlist->tail;
      } else {
        jlist->head = current->next;
        jlist->tail = current == jlist->tail ? NULL : jlist->tail;
      }
      free(current);
      break;
    }
    prev = current;
    current = current->next;
  }
}

void check_address(void *p) {
  if (p == NULL) {
    printf("Unable to allocate memory.\n");
    exit(EXIT_FAILURE);
  }
}
