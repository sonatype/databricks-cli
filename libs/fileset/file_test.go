package fileset

import (
	"path/filepath"
	"testing"

	"github.com/databricks/cli/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestNotebookFileIsNotebook(t *testing.T) {
	f := NewNotebookFile(nil, "", "")
	isNotebook, err := f.IsNotebook()
	require.NoError(t, err)
	require.True(t, isNotebook)
}

func TestSourceFileIsNotNotebook(t *testing.T) {
	f := NewSourceFile(nil, "", "")
	isNotebook, err := f.IsNotebook()
	require.NoError(t, err)
	require.False(t, isNotebook)
}

func TestUnknownFileDetectsNotebook(t *testing.T) {
	tmpDir := t.TempDir()

	t.Run("file", func(t *testing.T) {
		path := testutil.Touch(t, tmpDir, "test.py")
		f := NewFile(nil, path, filepath.Base(path))
		isNotebook, err := f.IsNotebook()
		require.NoError(t, err)
		require.False(t, isNotebook)
	})

	t.Run("notebook", func(t *testing.T) {
		path := testutil.TouchNotebook(t, tmpDir, "notebook.py")
		f := NewFile(nil, path, filepath.Base(path))
		isNotebook, err := f.IsNotebook()
		require.NoError(t, err)
		require.True(t, isNotebook)
	})
}
