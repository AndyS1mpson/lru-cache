package cache

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
 Название тестов:
 Test{структура}_{метод}_{данные}_{ожидание}
 Пример: TestBuilder_BuildFeneralStep_hasID_buildStep

*/

func TestLRUCache_Get_Item_OK(t *testing.T) {
	t.Parallel()
	// Arrange
	cache := NewLRUCache(5)
	key := "test_key"
	value := "test_value"
	elem := cache.queue.PushFront(value)
	cache.items[key] = elem

	// Act
	el, ok := cache.Get(key)

	// Assert
	require.True(t, ok)
	require.Equal(t, value, el)
}

func TestLRUCache_Set_Item_OK(t *testing.T) {
	t.Parallel()

	// Arrange
	cache := NewLRUCache(5)
	key := "test_key"
	value := "test_value"

	// Act
	err := cache.Set(key, value)

	// Assert
	el := cache.items[key]

	require.NoError(t, err)
	require.Equal(t, value, el)

}

func TestLRU_Set_ExistElementWithFULlQueueSync_MoveToFront(t *testing.T) {
	t.Parallel()
	// Arrange
	cache := NewLRUCache(3)
	cache.Set("Vasya", 10)
	cache.Set("Petya", "opa")
	cache.Set("Kolya", []int{1, 2})

	// Act
	cache.Set("Vasva", 15)

	resultFront, _ := cache.queue.Front().Value.(*list.Element)
	resultBack, _ := cache.queue.Back().Value.(*list.Element)

	require.Equal(t, resultFront, 15)
	require.Equal(t, resultBack, "opa")
	require.Equal(t, cache.queue.Len(), 3)

}

func TestLRUCache_Set_MoreThanCap_MaxCap(t *testing.T) {
	t.Parallel()

	// Arrange
	cache := NewLRUCache(5)

	cache.Set("test_1", "value_1")
	cache.Set("test_2", "value_2")
	cache.Set("test_3", "value_3")
	cache.Set("test_4", "value_4")
	cache.Set("test_5", "value_5")
	cache.Set("test_6", "value_6")

	// Assert
	el := cache.items["test_1"]

	require.Nil(t, el)
}

func TestLRUCache_Delete_Item_OK(t *testing.T) {
	t.Parallel()
	// Arrange
	cache := NewLRUCache(5)
	key := "test_key"
	value := "test_value"
	elem := cache.queue.PushFront(value)
	cache.items[key] = elem

	// Act
	cache.Delete(key)

	// Assert
	_, ok := cache.items[key]

	require.False(t, ok)
}

func TestLRUCache_Clear_Cache_OK(t *testing.T) {
	t.Parallel()
	// Arrange
	cache := NewLRUCache(5)
	key := "test_key"
	value := "test_value"
	elem := cache.queue.PushFront(value)
	cache.items[key] = elem

	// Act
	cache.Clear()

	// Assert
	el := cache.queue.Front()
	require.Zero(t, len(cache.items))
	require.Nil(t, el)
}

func TestLRUCache_Count_Items_OK(t *testing.T) {
	t.Parallel()
	// Arrange
	cache := NewLRUCache(5)
	key := "test_key"
	value := "test_value"
	elem := cache.queue.PushFront(value)
	cache.items[key] = elem

	// Act
	el := cache.Count()

	// Assert
	require.Equal(t, el, 1)
}
